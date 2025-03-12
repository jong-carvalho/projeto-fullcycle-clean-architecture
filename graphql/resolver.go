package graphql

import (
	"orders-service/models"
	"orders-service/services"
)

type Resolver struct{}

func (r *Resolver) ListOrders() []models.Order {
	return services.ListOrders()
}
