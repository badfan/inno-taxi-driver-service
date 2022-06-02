package models

import (
	"time"

	"github.com/google/uuid"
)

type TaxiType string

const (
	TaxiTypeEconomy  TaxiType = "economy"
	TaxiTypeComfort  TaxiType = "comfort"
	TaxiTypeBusiness TaxiType = "business"
	TaxiTypeElectro  TaxiType = "electro"
)

type Driver struct {
	ID           int32     `json:"id"`
	DriverUuid   uuid.UUID `json:"driver_uuid"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	TaxiType     TaxiType  `json:"taxi_type"`
	IsBusy       bool      `json:"is_busy"`
	DriverRating float32   `json:"driver_rating"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
