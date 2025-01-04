package server

import (
	"context"
	"playground/internal/logger"
	"playground/internal/models"
	"playground/internal/services/orders"
	"playground/internal/services/works"
	"playground/pkg/events"
)

func registerEventsListeners() {
	l := events.DefaultListeners
	l.AddListener(context.TODO(), "order:new", func(e events.Event[any]) {
		o := e.Payload.(orders.Order)
		ws, err := models.FindNearestWorkshop(o.Address.GeoLat, o.Address.GeoLon)
		if err != nil {
			logger.Error(err)
			return
		}
		wm := models.NewWorkModel()
		wm.Model = models.Work{
			OrderId:    o.Id,
			WorkshopId: ws.Id,
			Status:     uint8(works.StatusPending),
			StartAt:    nil,
		}
		_, err = wm.Create()
		if err != nil {
			logger.Error(err)
		}
	})
	// l.Listen() // not necessary when running from main function
}
