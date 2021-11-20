package model

import "fmt"

//RentObjectTypes for RentObjectType
const (
	House RentObjectType = "house"
	Car   RentObjectType = "car"
)

//RentObjectType of Rent entity
type RentObjectType string

func (r RentObjectType) String() string {
	return string(r)
}

//Icon of RentObjectType
func (r RentObjectType) Icon() string {
	switch r {
	case Car:
		return "🚗"
	case House:
		return "🏚"
	default:
		panic(fmt.Errorf("RentObjectType: error unexpected value %s", r))
	}
	return r.String()
}

//Full type of RentObjectType
func (r RentObjectType) Full() string {
	switch r {
	case Car:
		return "🚗 car"
	case House:
		return "🏚 house"
	default:
		panic(fmt.Errorf("RentObjectType: error unexpected value %s", r))
	}
	return r.String()
}
