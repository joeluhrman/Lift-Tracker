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

type apiError struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

func (e apiError) Error() string {
	return e.Err
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Err: "internal server"})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

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

	if !utils.PasswordMeetsRequirements(user.Password) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.IsAdmin = false

	err = db.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
