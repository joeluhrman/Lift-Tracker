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
		dbDriver     = "pgx"
		dbApiKeyPath = "./db/api_key.txt"

		serverPort = "3000"
	)

	var (
		dbApiKey = string(mustReadFile(dbApiKeyPath))
		dbURL    = db.NewURL("Lift-Tracker", "db.bit.io" /*dbPort,*/, "jaluhrman", dbApiKey)
	)

	db.MustInit(&db.Config{
		Driver: dbDriver,
		URL:    dbURL,
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
