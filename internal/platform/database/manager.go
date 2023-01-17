package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxConns = 10
)

func New(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxConns)
	return db, nil
}
