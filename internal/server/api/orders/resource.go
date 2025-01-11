package orders

import "playground/internal/services/orders"

type orderAddressResource struct {
	FullAddress string  `json:"fullAddress" example:"г Санкт-Петербург"`
	GeoLat      float32 `json:"geoLat" example:"59.939083"`
	GeoLon      float32 `json:"geoLon" example:"30.31588"`
}
type orderResource struct {
	Id      int                  `json:"id" example:"1000"`
	Time    string               `json:"time" example:"2028-01-01T13:00:00Z"`
	Address orderAddressResource `json:"address"`
	Comment string               `json:"comment" example:"My pc is broken, so I can't accomplish my Golang project!"`
}

func toResource(o orders.Order) orderResource {
	return orderResource{
		Id:      o.Id,
		Time:    o.Time,
		Comment: o.Comment,
		Address: orderAddressResource{
			FullAddress: o.Address.FullAddress,
			GeoLat:      o.Address.GeoLat,
			GeoLon:      o.Address.GeoLon,
		},
	}
}
