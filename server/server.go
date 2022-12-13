// Contains functionality for creating/starting a Server and handling HTTP requests.
package server

import (
	"context"
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
	endLogout    = "/logout"
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
		r.Post(endLogout, makeHTTPHandler(s.handleLogout))
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
	user := &types.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return newApiError(http.StatusBadRequest, err.Error())
	}

	if !storage.UsernameMeetsRequirements(user.Username) {
		return newApiError(http.StatusNotAcceptable, errors.New("username does not meet requirements").Error())
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

func setSession(s *types.Session, w http.ResponseWriter) {
	http.SetCookie(w, s.Cookie())
}

func getSessionToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(types.SessionKey)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (s *Server) middlewareAuthSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getSessionToken(r)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, err.Error())
			return
		}

		userID, err := s.storage.AuthenticateSession(token)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, err.Error())
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {
	user := &types.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return newApiError(http.StatusBadRequest, err.Error())
	}

	userID, err := s.storage.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		return newApiError(http.StatusUnauthorized, err.Error())
	}

	err = s.storage.DeleteSessionByUserID(userID)
	if err != nil {
		return newApiError(http.StatusInternalServerError, err.Error())
	}

	session := types.NewSession(userID)
	err = s.storage.InsertSession(session)
	if err != nil {
		return newApiError(http.StatusInternalServerError, err.Error())
	}

	setSession(session, w)

	return writeJSON(w, http.StatusOK, nil)
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) error {
	token, err := getSessionToken(r)
	if err != nil {
		return newApiError(http.StatusNotFound, err.Error())
	}

	http.SetCookie(w, &http.Cookie{
		Name:  types.SessionKey,
		Value: "",
	})

	err = s.storage.DeleteSessionByToken(token)
	if err != nil {
		return newApiError(http.StatusInternalServerError, err.Error())
	}

	return writeJSON(w, http.StatusOK, nil)
}
