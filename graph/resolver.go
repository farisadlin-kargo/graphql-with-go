package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/farisadlin-kargo/kargo-trucks/graph/model"
)

type Resolver struct {
	Trucks    []*model.Truck
	Shipments []*model.Shipment
}

func (r *Resolver) Init() {
	for i := 0; i < 20; i++ {
		truck := &model.Truck{
			ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
			PlateNo: fmt.Sprintf("B %d CD", len(r.Trucks)+1),
		}
		r.Trucks = append(r.Trucks, truck)
	}

	records := [][]string{
		{"id", "plate_no"},
		{"John", "Doe"},
		{"Lucy", "Smith"},
		{"Brian", "Bethamy"},
	}
	f, err := os.Create("trucks.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
