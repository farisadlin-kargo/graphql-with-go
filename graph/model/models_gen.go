// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EmailDataInput struct {
	Email string `json:"email"`
}

type Shipment struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	DeliveryDate string `json:"deliveryDate"`
	Truck        *Truck `json:"truck"`
}

type Truck struct {
	ID      string `json:"id"`
	PlateNo string `json:"plateNo"`
}
