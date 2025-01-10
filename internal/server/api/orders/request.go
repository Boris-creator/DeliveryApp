package orders

import (
	"playground/internal/services/addresses"
	"strconv"
)

type address struct {
	FullAddress string `json:"fullAddress"`
	GeoLat      string `json:"geo_lat"`
	GeoLon      string `json:"geo_lon"`
}

func (addr address) Address() addresses.Address {
	geoLat, _ := strconv.ParseFloat(addr.GeoLat, 32)
	geoLon, _ := strconv.ParseFloat(addr.GeoLon, 32)
	return addresses.Address{
		FullAddress: addr.FullAddress,
		GeoLat:      float32(geoLat),
		GeoLon:      float32(geoLon),
	}
}

type SaveOrderRequest struct {
	Address address `json:"address" validate:"required"`
	Time    string  `json:"time" validate:"required,datetime=2006-01-02T15:04"`
	Comment string  `json:"comment" validate:"omitempty,gte=2,lte=250"`
}
