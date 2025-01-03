package workshops

import "playground/internal/services/addresses"

type Workshop struct {
	ID        uint
	Name      string
	AddressID uint
	Address   addresses.Address
}
