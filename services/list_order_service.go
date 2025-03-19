package services

import (
	"projeto-fullcycle-clean-architecture/models"
	"projeto-fullcycle-clean-architecture/repositories"
)

func ListOrders() []models.Order {
	return repositories.GetOrders()
}
