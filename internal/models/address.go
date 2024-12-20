package models

type Address struct {
	FullAddress string `db:"full_address"`
}

type AddressModel = Model[Address]
