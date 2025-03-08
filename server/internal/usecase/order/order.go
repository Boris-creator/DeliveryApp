package order

import "playground.com/server/internal/usecase/address"

type Order struct {
	Time    string
	Address address.Address
	Comment string
	Id      int
}
