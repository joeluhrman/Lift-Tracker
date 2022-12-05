package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joeluhrman/Lift-Tracker/db"
	"github.com/joeluhrman/Lift-Tracker/utils"
)

const (
	routeApiV1   = "/api/v1"
	endCreateAcc = "create-account"
)

func setupEndpoints(r *chi.Mux) {
	r.Route(routeApiV1, func(r chi.Router) {
		r.Post(endCreateAcc, handleCreateAccount)
	})
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	user := &db.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.IsAdmin = false

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
