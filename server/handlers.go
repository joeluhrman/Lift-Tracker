package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	ROUTE_API_V1   = "/api/v1"
	END_CREATE_ACC = "create-account"
)

func setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/create-account", handleCreateAccount)
	})
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("placeholder")
}
