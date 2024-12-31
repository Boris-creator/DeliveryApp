package orders

import "playground/internal/services/orders"

type orderAddressResource struct {
	FullAddress string  `json:"fullAddress"`
	GeoLat      float32 `json:"geoLat"`
	GeoLon      float32 `json:"geoLon"`
}
type orderResource struct {
	Id      int                  `json:"id"`
	Time    string               `json:"time"`
	Address orderAddressResource `json:"address"`
	Comment string               `json:"comment"`
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
