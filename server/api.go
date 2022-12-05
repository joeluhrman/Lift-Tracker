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

func newApiError(status int, err string) apiError {
	return apiError{
		Status: status,
		Err:    err,
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, e.Err)
			} else {
				writeJSON(w, http.StatusInternalServerError, err)
			}
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
		r.Post(endCreateAcc, makeHTTPHandler(handleCreateAccount))
	})
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	user := &db.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return newApiError(http.StatusBadRequest, err.Error())
	}

	if !utils.PasswordMeetsRequirements(user.Password) {
		return newApiError(http.StatusNotAcceptable, err.Error())
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return newApiError(http.StatusInternalServerError, err.Error())
	}

	user.IsAdmin = false

	err = db.InsertUser(user)
	if err != nil {
		return newApiError(http.StatusConflict, err.Error())
	}

	return writeJSON(w, http.StatusAccepted, nil)
}
