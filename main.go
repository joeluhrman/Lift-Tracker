package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/joeluhrman/Lift-Tracker/db"
	"github.com/joeluhrman/Lift-Tracker/server"
)

func main() {
	const (
		serverPort = ":3000"
	)

	var (
		isProd   bool
		dbType   db.DBType
		dbDriver string
		dbPath   string
	)

	if len(os.Args) > 1 {
		isProd = os.Args[1] == "-prod"
	}

	if isProd {
		dbApiKey := string(mustReadFile("./db/api_key.txt"))
		url := db.NewPostgresqlURL("Lift-Tracker", "db.bit.io", "jaluhrman", dbApiKey)

		dbType = db.PGSQL
		dbDriver = "pgx"
		dbPath = url.String()

	} else {
		dbType = db.SQLite
		dbDriver = "sqlite3"
		dbPath = "./test.db"
	}

	db.MustInit(&db.Config{
		Type:   dbType,
		Driver: dbDriver,
		Path:   dbPath,
	})

	defer db.MustClose()

	server.MustStart(&server.Config{
		Port: serverPort,
		Middlewares: []func(http.Handler) http.Handler{
			middleware.Logger,
		},
	})
}

func mustReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}
