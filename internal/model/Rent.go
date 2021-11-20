package model

import (
	"github.com/shopspring/decimal"
	"time"
)

//Rent model
type Rent struct {
	ID         uint64
	RenterID   uint64          `json:"renter_id"`
	ObjectType RentObjectType  `json:"object_type"`
	ObjectInfo string          `json:"object_info"`
	Period     time.Duration   `json:"period"`
	Price      decimal.Decimal `json:"price"`
}
