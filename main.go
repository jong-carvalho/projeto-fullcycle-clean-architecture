package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	//"projeto-fullcycle-clean-architecture/proto"
	"time"

	"github.com/gin-gonic/gin"
)

// Modelo Order
type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	TotalAmount  float64   `json:"total_amount"`
	CreatedAt    time.Time `json:"created_at"`
}

// Configuração do banco de dados
var DB *gorm.DB
var err error

// Função para configurar e conectar ao banco de dados
func setupDatabase() {
	// Configuração de conexão com o banco
	dsn := "host=localhost user=user password=password dbname=orders_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Migrar o modelo para o banco de dados
	DB.AutoMigrate(&Order{})
}

// Função para inicializar o servidor gRPC
//func setupGRPC() {
//	// Cria o listener TCP
//	listener, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	// Cria o servidor gRPC
//	grpcServer := grpc.NewServer()
//
//	// Registra o serviço OrderServiceServer
//	orderService := &grpc.OrderServiceServer{}
//	proto.RegisterOrderServiceServer(grpcServer, orderService)
//
//	// Inicia o servidor gRPC
//	log.Println("gRPC Server is running on port :50051...")
//	if err := grpcServer.Serve(listener); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

func main() {
	// Configuração do banco de dados
	setupDatabase()

	// Configuração do servidor gRPC em goroutine para rodar simultaneamente
	//go setupGRPC()

	// Criando o roteador Gin
	r := gin.Default()

	// Endpoint para verificar se a API está rodando
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API is running!"})
	})

	// Endpoint para listar orders
	r.GET("/orders", func(c *gin.Context) {
		var orders []Order
		// Buscar os pedidos no banco de dados
		if err := DB.Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar pedidos"})
			return
		}

		c.JSON(http.StatusOK, orders)
	})

	// Endpoint para criar um novo pedido (POST)
	r.POST("/orders", func(c *gin.Context) {
		var order Order

		// Validar dados da requisição
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos"})
			return
		}

		// Definir a data de criação como o horário atual
		order.CreatedAt = time.Now()

		// Inserir o novo pedido no banco de dados
		if err := DB.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar pedido"})
			return
		}

		// Retornar o pedido criado com o ID gerado
		c.JSON(http.StatusCreated, order)
	})

	// Definir a porta do servidor HTTP
	port := ":8080"
	fmt.Println("Server running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
