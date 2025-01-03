package works

import (
	"playground/internal/services/orders"
	"playground/internal/services/workshops"
	"time"
)

type Status uint8

const (
	StatusPending Status = iota
	StatusInProgress
	StatusDone
)

type Work struct {
	WorkshopId int
	Workshop   workshops.Workshop
	OrderId    int
	Order      orders.Order
	StartAt    time.Time
	Status     Status
}
