package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	conn *sql.DB
)

type Config struct {
	Driver string
	Path   string
}

func Init(cfg *Config) error {
	if conn != nil {
		panic(errors.New("db cannot be initialized more than once"))
	}

	var err error
	conn, err = sql.Open(cfg.Driver, cfg.Path)

	return err
}

func Close() error {
	return conn.Close()
}
