package orders

import (
	"context"
	"fmt"

	"playground/internal/models"
)

func SaveOrder(ctx context.Context, order Order) (ord *Order, err error) {
	tx, err := models.Conn.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed saving order: %w", err)
		}
	}()
	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
	}()

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
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &Order{
		Id:      o.Id,
		Address: addr.ToAddress(),
		Time:    o.Time,
		Comment: o.Comment,
	}, err
}
