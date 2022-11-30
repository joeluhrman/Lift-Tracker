package db

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	driver = "sqlite3"
	path   = "../database.db"
)

var (
	conn *sql.DB
)

func Init() {
	if conn != nil {
		panic(errors.New("db cannot be initialized more than once"))
	}

	if err := setup(); err != nil {
		panic(err)
	}

	var err error
	conn, err = sql.Open(driver, path)
	if err != nil {
		panic(err)
	}
}

func setup() error {
	_, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	return err
}

func Close() {
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}
