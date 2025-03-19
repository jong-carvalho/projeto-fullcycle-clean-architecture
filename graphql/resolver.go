package graphql

import (
	"projeto-fullcycle-clean-architecture/models"
	"projeto-fullcycle-clean-architecture/services"
)

type Resolver struct{}

func (r *Resolver) ListOrders() []models.Order {
	return services.ListOrders()
}
