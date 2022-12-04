package db

import (
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

var (
	conn *sql.DB
)

type Config struct {
	Driver string
	Path   string
}

func MustConnect(cfg *Config) {
	if conn != nil {
		panic(errors.New("db cannot be initialized more than once"))
	}

	var err error
	conn, err = sql.Open(cfg.Driver, cfg.Path)
	if err != nil {
		panic(err)
	}
}

func MustClose() {
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}
