package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joeluhrman/Lift-Tracker/storage"
)

const (
	routeApiV1   = "/api/v1"
	endCreateAcc = "/create-account"
)

type middleware func(http.Handler) http.Handler

type Server struct {
	port        string
	router      *chi.Mux
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
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
