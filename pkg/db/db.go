package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=cpm dbname=cpm sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
