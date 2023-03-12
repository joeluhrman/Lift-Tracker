// Contains functionality for creating/starting a Server and handling HTTP requests.
package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joeluhrman/Lift-Tracker/storage"
	"github.com/joeluhrman/Lift-Tracker/types"
)

const (
	routeApiV1         = "/api/v1"
	endUser            = "/user"
	endLogin           = "/login"
	endLogout          = "/logout"
	endExerciseType    = "/exercise-Type"
	endWorkoutTemplate = "/workout-template"

	keyUserID  = "user_id"
	keySession = types.SessionKey
)

type Server struct {
	http.Server
	storage storage.Storage
}

func New(port string, storage storage.Storage, middlewares ...func(http.Handler) http.Handler) *Server {
	httpServer := http.Server{
		Addr: port,
	}
	s := &Server{
		Server:  httpServer,
		storage: storage,
	}

	router := chi.NewRouter()
	for i := range middlewares {
		router.Use(middlewares[i])
	}
	s.setupEndpoints(router)
	s.Handler = router

	return s
}

func (s *Server) MustStart() {
	log.Println("server running on port " + s.Addr)
	s.ListenAndServe()
}

func (s *Server) MustShutdown(shutdownCtx context.Context) {
	err := s.Shutdown(shutdownCtx)
	if err != nil {
		panic(err)
	}
}

func (s *Server) setupEndpoints(router *chi.Mux) {
	router.Route(routeApiV1, func(r chi.Router) {
		r.Post(endUser, s.handleCreateUser)
		r.Post(endLogin, s.handleLogin)
		r.Post(endLogout, s.handleLogout)

		r.Group(func(auth chi.Router) {
			auth.Use(s.middlewareAuthSession)

			auth.Get(endExerciseType, s.handleGetExerciseTypes)

			auth.Get(endWorkoutTemplate, s.handleGetWorkoutTemplates)
			auth.Post(endWorkoutTemplate, s.handleCreateWorkoutTemplate)
		})
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if !storage.UsernameMeetsRequirements(user.Username) {
		writeJSON(w, http.StatusNotAcceptable, "username does not meet requirements")
		return
	}

	if !storage.PasswordMeetsRequirements(user.Password) {
		writeJSON(w, http.StatusNotAcceptable, "password does not meet requirements")
		return
	}

	err = s.storage.CreateUser(user)
	if err != nil {
		writeJSON(w, http.StatusConflict, err.Error())
		return
	}

	writeJSON(w, http.StatusAccepted, nil)
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
			return
		}

		ctx := context.WithValue(r.Context(), keyUserID, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := s.storage.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	err = s.storage.DeleteSessionByUserID(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	session := types.NewSession(userID)
	err = s.storage.CreateSession(session)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	setSession(session, w)

	writeJSON(w, http.StatusOK, nil)
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	token, err := getSessionToken(r)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  keySession,
		Value: "",
	})

	err = s.storage.DeleteSessionByToken(token)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func (s *Server) handleGetExerciseTypes(w http.ResponseWriter, r *http.Request) {
	eTypes, err := s.storage.GetExerciseTypes()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusFound, eTypes)
}

func (s *Server) handleCreateWorkoutTemplate(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)

	wTemp := &types.WorkoutTemplate{}
	if err := json.NewDecoder(r.Body).Decode(wTemp); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	wTemp.UserID = userID

	if err := s.storage.CreateWorkoutTemplate(wTemp); err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
	}

	writeJSON(w, http.StatusAccepted, nil)
}

func (s *Server) handleGetWorkoutTemplates(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)

	wTemps, err := s.storage.GetWorkoutTemplates(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusFound, wTemps)
}
