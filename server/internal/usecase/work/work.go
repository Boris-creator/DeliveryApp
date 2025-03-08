package work

import (
	"time"

	"playground.com/server/internal/usecase/order"
	"playground.com/server/internal/usecase/workshop"
)

type Status uint8

const (
	StatusPending Status = iota
	StatusInProgress
	StatusDone
)

type Work struct {
	WorkshopId int
	Workshop   workshop.Workshop
	OrderId    int
	Order      order.Order
	StartAt    time.Time
	Status     Status
}
