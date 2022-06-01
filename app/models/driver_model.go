package models

import (
	"time"

	"github.com/google/uuid"
)

type TaxiTypes string

const (
	TaxiTypesEconomy  TaxiTypes = "economy"
	TaxiTypesComfort  TaxiTypes = "comfort"
	TaxiTypesBusiness TaxiTypes = "business"
	TaxiTypesElectro  TaxiTypes = "electro"
)

type Driver struct {
	ID           int32     `json:"id"`
	DriverUuid   uuid.UUID `json:"driver_uuid"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	TaxiType     TaxiTypes `json:"taxi_type"`
	IsBusy       bool      `json:"is_busy"`
	DriverRating float32   `json:"driver_rating"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
