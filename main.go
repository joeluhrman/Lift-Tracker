package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	pgStore := storage.NewPostgres(pgDriver, pgURL)
	pgStore.MustConnect()
	defer pgStore.MustClose()

	server := server.New(*listenaddr, pgStore, middleware.Logger)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		server.MustShutdown(shutdownCtx)
		log.Println("Server successfully shutdown")
		serverStopCtx()
	}()

	server.MustStart()

	<-serverCtx.Done()
}
