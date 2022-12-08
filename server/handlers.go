package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joeluhrman/Lift-Tracker/types"
)

func (s *Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	user := &types.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return newApiError(http.StatusBadRequest, err.Error())
	}

	if !PasswordMeetsRequirements(user.Password) {
		return newApiError(http.StatusNotAcceptable, errors.New("password does not meet requirements").Error())
	}

	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return newApiError(http.StatusInternalServerError, err.Error())
	}

	err = s.storage.InsertUser(user, false)
	if err != nil {
		return newApiError(http.StatusConflict, err.Error())
	}

	return writeJSON(w, http.StatusAccepted, nil)
}
