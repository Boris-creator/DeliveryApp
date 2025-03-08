package models

import "github.com/jmoiron/sqlx"

type Order struct {
	Id        int    `db:"id"         sql:"omit_on_insert"`
	AddressId int    `db:"address_id"`
	CreatedAt string `db:"created_at" sql:"omit_on_insert"`
	Time      string `db:"time"`
	Comment   string `db:"comment"`
}

type OrderModel = Model[Order]

func (o Order) ToOrderModel(db *sqlx.DB) OrderModel {
	return OrderModel{DB: db, Model: o, TableName: "orders"}
}
