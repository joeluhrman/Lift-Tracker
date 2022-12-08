package storage

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joeluhrman/Lift-Tracker/types"
)

const (
	tableUser = "users"
)

type PostgresStorage struct {
	conn   *sql.DB
	driver string
	url    string
}

// Call PostgresStorage.MustConnect() to initialize connection.
func NewPostgresStorage(driver string, url string) *PostgresStorage {
	return &PostgresStorage{
		driver: driver,
		url:    url,
	}
}

func (p *PostgresStorage) MustConnect() {
	var err error
	p.conn, err = sql.Open(p.driver, p.url)
	if err != nil {
		panic(err)
	}

	err = p.conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("connected to database %s", p.url)
}

func (p *PostgresStorage) MustClose() {
	err := p.conn.Close()
	if err != nil {
		panic(err)
	}
}

func (p *PostgresStorage) InsertUser(user *types.User, isAdmin bool) error {
	statement := "INSERT INTO " + tableUser + " (username, password, is_admin) VALUES ($1, $2, $3)"

	_, err := p.conn.Exec(statement, user.Username, user.Password, isAdmin)

	return err
}