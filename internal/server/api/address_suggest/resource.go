package address_suggest

type AddressData struct {
	GeoLat string `json:"geo_lat"`
	GeoLon string `json:"geo_lon"`
}
type resource struct {
	Value string      `json:"fullAddress"`
	Data  AddressData `json:"data"`
}
