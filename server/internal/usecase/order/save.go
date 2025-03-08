package order

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"playground.com/server/internal/models"
)

func SaveOrder(ctx context.Context, db sqlx.DB, order Order) (ord *Order, err error) {
	tx, err := db.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			err = fmt.Errorf("saving order: %w", err)
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

	a := models.FromAddress(order.Address).ToAddressModel(&db)

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
	mo := mOrder.ToOrderModel(&db)
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
