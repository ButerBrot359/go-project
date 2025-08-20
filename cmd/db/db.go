package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
	var err error

	defer DB.Close()

	DB, err = sql.Open("postgres", "user=buterbrot dbname=buterdb sslmode=disable")

	return err
}