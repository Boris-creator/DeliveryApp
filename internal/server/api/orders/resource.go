package orders

import "playground/internal/services/orders"

type orderAddressResource struct {
	FullAddress string  `example:"г Санкт-Петербург"                json:"fullAddress"`
	GeoLat      float32 `example:"59.939083"                        json:"geoLat"`
	GeoLon      float32 `example:"30.31588"                         json:"geoLon"`
}
type orderResource struct {
	Id      int                  `example:"1000"                                                      json:"id"`
	Time    string               `example:"2028-01-01T13:00:00Z"                                      json:"time"`
	Address orderAddressResource `json:"address"`
	Comment string               `example:"My pc is broken, so I can't accomplish my Golang project!" json:"comment"`
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
