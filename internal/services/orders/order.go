package orders

import "playground/internal/services/addresses"

type Order struct {
	Time    string
	Address addresses.Address
	Comment string
	Id      int
}
