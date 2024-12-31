package models

import "playground/internal/services/addresses"

type Address struct {
	FullAddress string  `db:"full_address"`
	GeoLat      float32 `db:"geo_lat"`
	GeoLon      float32 `db:"geo_lon"`
	Id          int     `db:"id" sql:"autoincrement"`
}

type AddressModel = Model[Address]

func (a Address) ToAddressModel() AddressModel {
	return AddressModel{DB: Conn, Model: a, TableName: "address"}
}

func (address Address) ToAddress() addresses.Address {
	return addresses.Address{
		FullAddress: address.FullAddress,
		GeoLat:      address.GeoLat,
		GeoLon:      address.GeoLon,
	}
}

func FromAddress(address addresses.Address) Address {
	return Address{
		FullAddress: address.FullAddress,
		GeoLat:      address.GeoLat,
		GeoLon:      address.GeoLon,
	}
}
