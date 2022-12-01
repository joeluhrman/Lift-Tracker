package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	server *Server
)

type Config struct {
	Port        string
	Middlewares []func(http.Handler) http.Handler
}

type Server struct {
	router *chi.Mux
	port   string
}

func (server *Server) useMiddlewares(middlewares []func(http.Handler) http.Handler) {
	for i := range middlewares {
		server.router.Use(middlewares[i])
	}
}

func newServer(cfg *Config) *Server {
	server := &Server{
		router: chi.NewRouter(),
		port:   cfg.Port,
	}

	server.useMiddlewares(cfg.Middlewares)

	return server
}

func MustStart(cfg *Config) {
	if server != nil {
		panic(errors.New("server cannot be started more than once"))
	}

	server = newServer(cfg)

	http.ListenAndServe(":"+server.port, server.router)
}
