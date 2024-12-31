package orders

import (
	"context"
	"playground/internal/models"
)

func SaveOrder(ctx context.Context, order Order) (*Order, error) {
	tx, err := models.Conn.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	a := models.FromAddress(order.Address).ToAddressModel()
	addr, err := a.FirstOrCreate(
		"WHERE full_address = $1",
		[]any{order.Address.FullAddress},
	)
	if err != nil {
		return nil, err
	}

	mOrder := models.Order{
		AddressId: addr.Id,
		Time:      order.Time,
		Comment:   order.Comment,
	}
	mo := mOrder.ToOrderModel()
	o, err := mo.Create()
	tx.Commit()
	return &Order{
		Id:      o.Id,
		Address: addr.ToAddress(),
		Time:    o.Time,
		Comment: o.Comment,
	}, err
}
