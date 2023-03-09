package main

import (
	"flag"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/joeluhrman/Lift-Tracker/server"
	"github.com/joeluhrman/Lift-Tracker/storage"
)

func main() {
	var (
		pgDriver = "pgx"
		pgApiKey = string(storage.MustReadFile("./api_keys/api_key_test.txt"))
		pgURL    = "postgresql://jaluhrman:" + pgApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test"

		listenaddr = flag.String("listenaddr", ":3000", "the port the server should listen on")
	)
	flag.Parse()

	pgStore := storage.NewPostgresStorage(pgDriver, pgURL)
	pgStore.MustConnect()
	defer pgStore.MustClose()

	server := server.New(*listenaddr, pgStore, middleware.Logger)
	server.MustStart()
}
