package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/joeluhrman/Lift-Tracker/db"
	"github.com/joeluhrman/Lift-Tracker/server"
)

var (
	prodDBApiKey = string(mustReadFile("./db/api_key.txt"))
	prodDBURL    = db.NewPostgresqlURL(
		"Lift-Tracker",
		"db.bit.io",
		"",
		"jaluhrman",
		prodDBApiKey)

	testDBApiKey = string(mustReadFile("./db/api_key_test.txt"))
	testDBURL    = db.NewPostgresqlURL(
		"Lift-Tracker-Test",
		"db.bit.io",
		"",
		"jaluhrman",
		testDBApiKey,
	)
)

func main() {
	const (
		serverPort = ":3000"
	)

	var (
		isProd   bool
		dbPath   string
		dbDriver = "pgx"
	)

	if len(os.Args) > 1 {
		isProd = os.Args[1] == "-prod"
	}

	if isProd {
		dbPath = prodDBURL.String()

	} else {
		dbPath = testDBURL.String()
	}

	db.MustConnect(&db.Config{
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
