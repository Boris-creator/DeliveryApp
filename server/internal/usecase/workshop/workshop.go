package workshop

import "playground.com/server/internal/usecase/address"

type Workshop struct {
	ID        uint
	Name      string
	AddressID uint
	Address   address.Address
}
