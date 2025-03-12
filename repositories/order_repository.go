package repositories

import (
	"orders-service/config"
	"orders-service/models"
)

func GetOrders() []models.Order {
	var orders []models.Order
	config.DB.Find(&orders)
	return orders
}
