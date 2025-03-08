package models

import (
	"github.com/jmoiron/sqlx"
	"playground.com/server/internal/usecase/address"
)

type Address struct {
	FullAddress string  `db:"full_address"`
	GeoLat      float32 `db:"geo_lat"`
	GeoLon      float32 `db:"geo_lon"`
	Id          int     `db:"id"           sql:"omit_on_insert"`
}

type AddressModel = Model[Address]

func (addr Address) ToAddressModel(db *sqlx.DB) AddressModel {
	return AddressModel{DB: db, Model: addr, TableName: "address"}
}

func (addr Address) ToAddress() address.Address {
	return address.Address{
		FullAddress: addr.FullAddress,
		GeoLat:      addr.GeoLat,
		GeoLon:      addr.GeoLon,
	}
}

func FromAddress(addr address.Address) Address {
	return Address{
		FullAddress: addr.FullAddress,
		GeoLat:      addr.GeoLat,
		GeoLon:      addr.GeoLon,
	}
}
