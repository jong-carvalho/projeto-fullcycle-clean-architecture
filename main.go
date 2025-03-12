package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Criando o roteador Gin
	r := gin.Default()

	// Endpoint para verificar se a API está rodando
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API is running!"})
	})

	// Endpoint para listar orders
	r.GET("/orders", func(c *gin.Context) {
		// Exemplo de pedidos (substitua por integração real com o banco de dados)
		orders := []map[string]interface{}{
			{"id": 1, "customer": "Jonatas Carvalho", "total": 99.99},
			{"id": 2, "customer": "Daniela Oliveira", "total": 49.50},
		}

		c.JSON(http.StatusOK, orders)
	})

	// Definir a porta do servidor
	port := ":8080"
	fmt.Println("Server running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
