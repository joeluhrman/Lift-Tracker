package db

import (
	"database/sql"
	"errors"
	"log"

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

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("connected to database %s", cfg.Path)
}

func MustClose() {
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}
