package model

//EventType of RentEvent
type EventType uint8

//EventStatus of RentEvent
type EventStatus uint8

//EventTypes for EventType
const (
	Created EventType = iota
	Updated
	Removed
)

//EventStatuses for EventStatus
const (
	Deferred EventStatus = iota
	Processed
)

//RentEvent with Rent entity
type RentEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Rent
}
