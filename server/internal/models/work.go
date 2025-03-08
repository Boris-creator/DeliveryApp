package models

import "github.com/jmoiron/sqlx"

type Work struct {
	WorkshopId int     `db:"workshop_id"`
	OrderId    int     `db:"order_id"`
	StartAt    *string `db:"start_at"`
	Status     uint8   `db:"status"`
}

type WorkModel = Model[Work]

func NewWorkModel(db *sqlx.DB) WorkModel {
	return WorkModel{DB: db, Model: Work{}, TableName: "works"}
}
