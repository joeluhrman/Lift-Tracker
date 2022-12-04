package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/joeluhrman/Lift-Tracker/db"
	"github.com/joeluhrman/Lift-Tracker/server"
)

var (
	testDBApiKey = string(mustReadFile("./db/api_key_test.txt"))

	TestDBConfig = &db.Config{
		Driver: "pgx",
		Path:   "postgresql://jaluhrman:" + testDBApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test",
	}

	TestServerConfig = &server.Config{
		Port: ":3000",
		Middlewares: []func(http.Handler) http.Handler{
			middleware.Logger,
		},
	}
)

func main() {
	var (
		isProd       bool
		dbConfig     *db.Config
		serverConfig *server.Config
	)

	if len(os.Args) > 1 {
		isProd = os.Args[1] == "-prod"
	}

	if isProd {
		dbConfig = TestDBConfig
		serverConfig = TestServerConfig
	} else {
		dbConfig = TestDBConfig
		serverConfig = TestServerConfig
	}

	db.MustConnect(dbConfig)
	defer db.MustClose()

	server.MustStart(serverConfig)
}

func mustReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}
