package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/joeluhrman/Lift-Tracker/db"
	"github.com/joeluhrman/Lift-Tracker/server"
	"github.com/joeluhrman/Lift-Tracker/utils"
)

func main() {
	var (
		dbApiKey = string(utils.MustReadFile("./api_keys/api_key_test.txt"))
		dbConfig = &db.Config{
			Driver: "pgx",
			Path:   "postgresql://jaluhrman:" + dbApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test",
		}

		serverConfig = &server.Config{
			Port: ":3000",
			Middlewares: []func(http.Handler) http.Handler{
				middleware.Logger,
			},
		}
	)

	db.MustConnect(dbConfig)
	defer db.MustClose()

	server.MustStart(serverConfig)
}
