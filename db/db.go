package db

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	conn *sql.DB
)

type Config struct {
	Driver string
	Path   string
}

func MustInit(cfg *Config) {
	if conn != nil {
		panic(errors.New("db cannot be initialized more than once"))
	}

	_, err := os.Stat(cfg.Path)
	if err != nil {
		file, err := os.Create(cfg.Path)
		if err != nil {
			panic(err)
		}

		file.Close()
	}

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
