package main

import (
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	ordersgrpc "projeto-fullcycle-clean-architecture/grpc"
	"projeto-fullcycle-clean-architecture/proto"
)

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	TotalAmount  float64   `json:"total_amount"`
	CreatedAt    time.Time `json:"created_at"`
}

var DB *gorm.DB
var err error

func setupDatabase() {

	dsn := "host=localhost user=user password=password dbname=orders_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	DB.AutoMigrate(&Order{})
}

func setupGRPC() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterOrderServiceServer(grpcServer, &ordersgrpc.OrderServiceServer{})

	reflection.Register(grpcServer)

	log.Println("gRPC Server is running on port :50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {

	setupDatabase()

	go setupGRPC()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API is running!"})
	})

	r.GET("/orders", func(c *gin.Context) {
		var orders []Order

		if err := DB.Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar pedidos"})
			return
		}

		c.JSON(http.StatusOK, orders)
	})

	r.POST("/orders", func(c *gin.Context) {
		var order Order

		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos"})
			return
		}

		order.CreatedAt = time.Now()

		if err := DB.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar pedido"})
			return
		}

		c.JSON(http.StatusCreated, order)
	})

	port := ":8080"
	fmt.Println("Server running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
