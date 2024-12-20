package models

import (
	//"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
	"playground/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var Conn *sqlx.DB

func Connect(cfg config.AppConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbServer))
	Conn = conn
	return conn, err
}

type Model[M any] struct {
	*sqlx.DB
	TableName string
	Model     M
}

func (m Model[M]) Exec(query string) error {
	_, err := m.NamedExec(query, m.Model)
	return err
}
