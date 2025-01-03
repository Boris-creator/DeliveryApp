package server

import (
	"context"
	"log"
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
			log.Println(err)
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
			log.Println(err)
		}
	})
	// l.Listen() // not necessary when running from main function
}
