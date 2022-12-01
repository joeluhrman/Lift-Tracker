package db

import (
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	conn *sql.DB
)

type URL struct {
	DBName   string
	Host     string
	Port     string
	Username string
	APIKey   string
}

func (url *URL) String() string {
	return "postgresql://" + url.Username + ":" + url.APIKey +
		"@" + url.Host + "/" + url.Username + "/" + url.DBName
}

func NewURL(dbName string, host string /*port string,*/, username string, apiKey string) *URL {
	return &URL{
		DBName: dbName,
		Host:   host,
		//Port:     port,
		Username: username,
		APIKey:   apiKey,
	}
}

type Config struct {
	Driver string
	URL    *URL
}

func MustInit(cfg *Config) {
	if conn != nil {
		panic(errors.New("db cannot be initialized more than once"))
	}

	var err error
	conn, err = sql.Open(cfg.Driver, cfg.URL.String())
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
