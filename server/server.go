package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	router *chi.Mux

	TestServerConfig = &Config{
		Port: ":3000",
		Middlewares: []func(http.Handler) http.Handler{
			middleware.Logger,
		},
	}
)

type Config struct {
	Port        string
	Middlewares []func(http.Handler) http.Handler
}

func MustStart(cfg *Config) {
	if router != nil {
		panic(errors.New("server cannot be started more than once"))
	}

	router := chi.NewRouter()

	useMiddlewares(router, cfg.Middlewares)
	setupEndpoints(router)

	http.ListenAndServe(cfg.Port, router)
}

func useMiddlewares(r *chi.Mux, middlewares []func(http.Handler) http.Handler) {
	for i := range middlewares {
		r.Use(middlewares[i])
	}
}
