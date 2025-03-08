package models

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"reflect"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"playground.com/server/internal/config"
)

func Connect(cfg config.AppConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(
		"pgx",
		fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.DbUser, cfg.DbPassword, net.JoinHostPort(cfg.DbHost, cfg.DbPort), cfg.DbServer),
	)

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

// First returns the first row matching the query, or an error if the query
// returns no rows or an error occurs.
//
// If cols set is not empty, it will be used as the column list in the SELECT
// statement. Otherwise, First will select all columns.
//
// The function returns the first row and an error if any.
func (m Model[M]) First(query string, args []any, cols []string) (M, error) {
	columns := "*"
	if len(cols) != 0 {
		columns = fmt.Sprintf("(%s)", strings.Join(cols, ", "))
	}

	q := fmt.Sprintf("SELECT %s FROM %s %s LIMIT 1", columns, m.TableName, query)

	var s M
	err := m.Get(&s, q, args...)

	return s, err
}

// Create inserts the record into the db table and returns the inserted row.
//
// The function returns the inserted row and an error if any.
func (m Model[M]) Create() (M, error) {
	var result M

	cols := getColumnNames(m.Model)

	values := make([]string, 0, len(cols))
	for _, c := range cols {
		values = append(values, ":"+c)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING *",
		m.TableName, strings.Join(cols, ", "), strings.Join(values, ", "),
	)
	q, args, _ := m.BindNamed(query, m.Model)
	err := m.QueryRowx(q, args...).StructScan(&result)

	return result, err
}

// FirstOrCreate returns the first record matching the query, or creates a new record if the query returns no rows.
//
// The function returns the first or just created row and an error if any.
func (m Model[M]) FirstOrCreate(query string, args []any) (M, error) {
	r, err := m.First(query, args, []string{})
	if errors.Is(err, sql.ErrNoRows) {
		return m.Create()
	}

	return r, err
}

func getColumnNames(s any) []string {
	var columns []string

	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return []string{}
	}

	for i := range v.NumField() {
		f := v.Type().Field(i)
		if !f.IsExported() {
			continue
		}

		key, ok := f.Tag.Lookup("db")
		sqlKey, _ := f.Tag.Lookup("sql")

		if !ok {
			key = f.Name
		}

		if key != "" && key != "-" && sqlKey != "omit_on_insert" {
			columns = append(columns, key)
		}
	}

	return columns
}
