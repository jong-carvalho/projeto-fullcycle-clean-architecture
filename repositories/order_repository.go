package repositories

import (
	"projeto-fullcycle-clean-architecture/config"
	"projeto-fullcycle-clean-architecture/models"
)

func GetOrders() []models.Order {
	var orders []models.Order
	config.DB.Find(&orders)
	return orders
}
