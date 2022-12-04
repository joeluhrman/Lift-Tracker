package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joeluhrman/Lift-Tracker/db"
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
	user := &db.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bleh"))
	}
}
