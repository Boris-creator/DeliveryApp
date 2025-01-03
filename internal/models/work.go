package models

type Work struct {
	WorkshopId int     `db:"workshop_id"`
	OrderId    int     `db:"order_id"`
	StartAt    *string `db:"start_at"`
	Status     uint8   `db:"status"`
}

type WorkModel = Model[Work]

func NewWorkModel() WorkModel {
	return WorkModel{DB: Conn, Model: Work{}, TableName: "works"}
}
