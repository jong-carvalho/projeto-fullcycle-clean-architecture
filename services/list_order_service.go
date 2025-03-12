package services

import (
	"orders-service/models"
	"orders-service/repositories"
)

func ListOrders() []models.Order {
	return repositories.GetOrders()
}
