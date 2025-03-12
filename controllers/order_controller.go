package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orders-service/services"
)

func GetOrders(c *gin.Context) {
	orders := services.ListOrders()
	c.JSON(http.StatusOK, orders)
}
