package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	tableUser = "users"
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

func MustReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func NewUser(username string, password string, isAdmin bool) *User {
	return &User{
		Username: username,
		Password: password,
		IsAdmin:  isAdmin,
	}
}

func InsertUser(user *User) error {
	statement := "INSERT INTO " + tableUser + " (username, password, is_admin) VALUES ($1, $2, $3)"

	_, err := conn.Exec(statement, user.Username, user.Password, user.IsAdmin)

	return err
}
