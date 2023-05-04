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
	endExerciseType    = "/exercise-type"
	endWorkoutTemplate = "/workout-template"
	endWorkoutLog      = "/workout-log"

	keyUserID  = "user_id"
	keySession = types.SessionKey
)

type server struct {
	http.Server
	storage storage.Storage
}

func New(port string, storage storage.Storage, middlewares ...func(http.Handler) http.Handler) *server {
	httpServer := http.Server{
		Addr: port,
	}
	s := &server{
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

func (s *server) MustStart() {
	log.Println("server running on port " + s.Addr)
	s.ListenAndServe()
}

func (s *server) MustShutdown(shutdownCtx context.Context) {
	err := s.Shutdown(shutdownCtx)
	if err != nil {
		panic(err)
	}

	log.Println("server successfully shutdown")
}

func (s *server) setupEndpoints(router *chi.Mux) {
	router.Route(routeApiV1, func(r chi.Router) {
		r.Post(endUser, s.handleCreateUser)
		r.Post(endLogin, s.handleLogin)

		r.Group(func(auth chi.Router) {
			auth.Use(s.middlewareAuthSession)
			//auth.Get(endLogin, s.handleIsLoggedIn)
			auth.Get(endUser, s.handleGetUser)
			auth.Get(endExerciseType, s.handleGetExerciseTypes)
			auth.Get(endWorkoutTemplate, s.handleGetWorkoutTemplates)
			auth.Get(endWorkoutLog, s.handleGetWorkoutLogs)
			auth.Post(endWorkoutTemplate, s.handleCreateWorkoutTemplate)
			auth.Post(endWorkoutLog, s.handleCreateWorkoutLog)
			auth.Post(endLogout, s.handleLogout)
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

func setSession(s types.Session, w http.ResponseWriter) {
	http.SetCookie(w, s.Cookie())
}

func getSessionToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(types.SessionKey)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (s *server) middlewareAuthSession(next http.Handler) http.Handler {
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

type credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	credentials := &credentials{}

	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if !storage.UsernameMeetsRequirements(credentials.Username) {
		writeJSON(w, http.StatusNotAcceptable, "username does not meet requirements")
		return
	}

	if !storage.PasswordMeetsRequirements(credentials.Password) {
		writeJSON(w, http.StatusNotAcceptable, "password does not meet requirements")
		return
	}

	user := &types.User{
		Username: credentials.Username,
		Email:    credentials.Email,
	}

	err = s.storage.CreateUser(user, credentials.Password)
	if err != nil {
		writeJSON(w, http.StatusConflict, err.Error())
		return
	}

	writeJSON(w, http.StatusAccepted, nil)
}

func (s *server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)
	user, err := s.storage.GetUser(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	credentials := &credentials{}

	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := s.storage.AuthenticateUser(credentials.Username, credentials.Password)
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

func (s *server) handleLogout(w http.ResponseWriter, r *http.Request) {
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

func (s *server) handleGetExerciseTypes(w http.ResponseWriter, r *http.Request) {
	eTypes, err := s.storage.GetExerciseTypes()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, eTypes)
}

func (s *server) handleCreateWorkoutTemplate(w http.ResponseWriter, r *http.Request) {
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

	writeJSON(w, http.StatusCreated, nil)
}

func (s *server) handleGetWorkoutTemplates(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)

	wTemps, err := s.storage.GetWorkoutTemplates(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, wTemps)
}

func (s *server) handleCreateWorkoutLog(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)

	wLog := &types.WorkoutLog{}
	if err := json.NewDecoder(r.Body).Decode(wLog); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	wLog.UserID = userID

	if err := s.storage.CreateWorkoutLog(wLog); err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (s *server) handleGetWorkoutLogs(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)

	wLogs, err := s.storage.GetWorkoutLogs(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusFound, wLogs)
}
