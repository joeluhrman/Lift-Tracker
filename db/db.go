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

type PostgresqlURL struct {
	DBName   string
	Host     string
	Port     string
	Username string
	APIKey   string
}

func (url *PostgresqlURL) String() string {
	return "postgresql://" + url.Username + ":" + url.APIKey +
		"@" + url.Host + url.Port + "/" + url.Username + "/" + url.DBName
}

func NewPostgresqlURL(dbName string, host string, port string, username string, apiKey string) *PostgresqlURL {
	return &PostgresqlURL{
		DBName:   dbName,
		Host:     host,
		Port:     port,
		Username: username,
		APIKey:   apiKey,
	}
}
