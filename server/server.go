package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	router *chi.Mux
)

type Config struct {
	Port        string
	Middlewares []func(http.Handler) http.Handler
}

func useMiddlewares(r *chi.Mux, middlewares []func(http.Handler) http.Handler) {
	for i := range middlewares {
		r.Use(middlewares[i])
	}
}

func MustStart(cfg *Config) {
	if router != nil {
		panic(errors.New("server cannot be started more than once"))
	}

	router := chi.NewRouter()

	useMiddlewares(router, cfg.Middlewares)

	http.ListenAndServe(cfg.Port, router)
}
