package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	CustomerName string  `json:"customer_name"`
	TotalAmount  float64 `json:"total_amount"`
	CreatedAt    string  `json:"created_at"`
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

func main() {

	setupDatabase()

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

	port := ":8080"
	fmt.Println("Server running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
