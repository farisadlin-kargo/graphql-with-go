package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/farisadlin-kargo/kargo-trucks/graph/model"
)

type Resolver struct {
	Trucks    []*model.Truck
	Shipments []*model.Shipment
}
