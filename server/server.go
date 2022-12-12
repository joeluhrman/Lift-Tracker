// Contains functionality for creating/starting a Server and handling HTTP requests.
package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joeluhrman/Lift-Tracker/storage"
	"github.com/joeluhrman/Lift-Tracker/types"
)

const (
	routeApiV1   = "/api/v1"
	endCreateAcc = "/user"
	endLogin     = "/login"

	errCodeBadJSON = http.StatusBadRequest
)

type middleware func(http.Handler) http.Handler

type Server struct {
	router      *chi.Mux
	port        string
	middlewares []middleware
	storage     storage.Storage
}

func New(port string, storage storage.Storage, middlewares ...middleware) *Server {
	return &Server{
		port:        port,
		middlewares: middlewares,
		storage:     storage,
	}
}

func (s *Server) MustStart() {
	s.router = chi.NewRouter()

	for i := range s.middlewares {
		s.router.Use(s.middlewares[i])
	}

	s.setupEndpoints()

	log.Println("server running on port " + s.port)
	http.ListenAndServe(s.port, s.router)
}

func (s *Server) setupEndpoints() {
	s.router.Route(routeApiV1, func(r chi.Router) {
		r.Post(endCreateAcc, makeHTTPHandler(s.handleCreateAccount))
		r.Post(endLogin, makeHTTPHandler(s.handleLogin))
	})
}

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
				return
			}
			writeJSON(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	var user *types.User

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return newApiError(errCodeBadJSON, err.Error())
	}

	if !storage.PasswordMeetsRequirements(user.Password) {
		return newApiError(http.StatusNotAcceptable, errors.New("password does not meet requirements").Error())
	}

	user.HashedPassword, err = storage.HashPassword(user.Password)
	if err != nil {
		return newApiError(http.StatusInternalServerError, err.Error())
	}

	err = s.storage.InsertUser(user, false)
	if err != nil {
		return newApiError(http.StatusConflict, err.Error())
	}

	return writeJSON(w, http.StatusAccepted, nil)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {
	var user *types.User

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return newApiError(errCodeBadJSON, err.Error())
	}

	userID, err := s.storage.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		return newApiError(http.StatusUnauthorized, err.Error())
	}

	session := types.NewSession(userID)
	err = s.storage.InsertSession(session)
	if err != nil {
		return newApiError(http.StatusInternalServerError, err.Error())
	}

	http.SetCookie(w, session.Cookie())

	return writeJSON(w, http.StatusOK, nil)
}
