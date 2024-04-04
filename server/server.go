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
)

type key string

const (
	keyUserID  key = "user_id"
	keySession key = types.SessionKey
)

type apiError struct {
	status int
	msg    string
}

func (a apiError) Error() string {
	return a.msg
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func executeAPIFunc(f apiFunc, w http.ResponseWriter, r *http.Request) error {
	err := f(w, r)
	if err != nil {
		a, ok := err.(apiError)
		if ok {
			writeJSON(w, a.status, a.msg)
		} else {
			writeJSON(w, http.StatusInternalServerError, err.Error())
		}
		return err
	}
	return nil
}

func makeHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		executeAPIFunc(f, w, r)
	}
}

func makeMiddleware(f apiFunc) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := executeAPIFunc(f, w, r)
			if err == nil {
				next.ServeHTTP(w, r)
			}
		})
	}
}

type server struct {
	http.Server
	storage storage.Storage
}

func (s *server) setupHandler(middlewares []func(http.Handler) http.Handler) {
	router := chi.NewRouter()
	for i := range middlewares {
		router.Use(middlewares[i])
	}

	router.Route(routeApiV1, func(r chi.Router) {
		r.Post(endUser, makeHandler(s.handleCreateUser))
		r.Post(endLogin, makeHandler(s.handleLogin))

		r.Group(func(auth chi.Router) {
			auth.Use(makeMiddleware(s.handleAuthSession))
			auth.Get(endUser, makeHandler(s.handleGetUser))
			auth.Get(endExerciseType, makeHandler(s.handleGetExerciseTypes))
			auth.Get(endWorkoutTemplate, makeHandler(s.handleGetWorkoutTemplates))
			auth.Get(endWorkoutLog, makeHandler(s.handleGetWorkoutLogs))
			auth.Post(endWorkoutTemplate, makeHandler(s.handleCreateWorkoutTemplate))
			auth.Post(endWorkoutLog, makeHandler(s.handleCreateWorkoutLog))
			auth.Post(endLogout, makeHandler(s.handleLogout))
		})
	})

	s.Handler = router
}

func New(port string, storage storage.Storage, middlewares ...func(http.Handler) http.Handler) *server {
	httpServer := http.Server{
		Addr: port,
	}
	s := &server{
		Server:  httpServer,
		storage: storage,
	}
	s.setupHandler(middlewares)

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

func getSessionToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(types.SessionKey)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (s *server) handleAuthSession(w http.ResponseWriter, r *http.Request) error {
	token, err := getSessionToken(r)
	if err != nil {
		return apiError{http.StatusUnauthorized, err.Error()}
	}

	userID, err := s.storage.AuthenticateSession(token)
	if err != nil {
		return apiError{http.StatusUnauthorized, err.Error()}
	}

	ctx := context.WithValue(r.Context(), keyUserID, userID)
	*r = *r.Clone(ctx)
	return nil
}

type credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *server) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	credentials := &credentials{}

	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		return apiError{http.StatusBadRequest, err.Error()}
	}

	user := &types.User{
		Username: credentials.Username,
		Email:    credentials.Email,
	}

	err = s.storage.CreateUser(user, credentials.Password)
	if err != nil {
		return apiError{http.StatusBadRequest, err.Error()}
	}

	return writeJSON(w, http.StatusAccepted, nil)
}

func setSession(s types.Session, w http.ResponseWriter) {
	http.SetCookie(w, s.Cookie())
}

func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) error {
	credentials := &credentials{}

	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		return apiError{http.StatusBadRequest, err.Error()}
	}

	userID, err := s.storage.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		return apiError{http.StatusUnauthorized, err.Error()}
	}

	err = s.storage.DeleteSessionByUserID(userID)
	if err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	session := types.NewSession(userID)
	err = s.storage.CreateSession(session)
	if err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	setSession(session, w)

	return writeJSON(w, http.StatusOK, nil)
}

func (s *server) handleLogout(w http.ResponseWriter, r *http.Request) error {
	token, err := getSessionToken(r)
	if err != nil {
		return apiError{http.StatusNotFound, err.Error()}
	}

	http.SetCookie(w, &http.Cookie{
		Name:  string(keySession),
		Value: "",
	})

	err = s.storage.DeleteSessionByToken(token)
	if err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	return writeJSON(w, http.StatusOK, nil)
}

func (s *server) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(keyUserID).(uint)
	user, err := s.storage.GetUser(userID)
	if err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *server) handleGetExerciseTypes(w http.ResponseWriter, r *http.Request) error {
	eTypes, err := s.storage.GetExerciseTypes()
	if err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	return writeJSON(w, http.StatusOK, eTypes)
}

func (s *server) handleCreateWorkoutTemplate(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(keyUserID).(uint)

	wTemp := &types.WorkoutTemplate{}
	if err := json.NewDecoder(r.Body).Decode(wTemp); err != nil {
		return apiError{http.StatusBadRequest, err.Error()}
	}

	wTemp.UserID = userID

	if err := s.storage.CreateWorkoutTemplate(wTemp); err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	return writeJSON(w, http.StatusCreated, nil)
}

func (s *server) handleGetWorkoutTemplates(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(keyUserID).(uint)

	wTemps, err := s.storage.GetWorkoutTemplates(userID)
	if err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	return writeJSON(w, http.StatusOK, wTemps)
}

func (s *server) handleCreateWorkoutLog(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(keyUserID).(uint)

	wLog := &types.WorkoutLog{}
	if err := json.NewDecoder(r.Body).Decode(wLog); err != nil {
		return apiError{http.StatusBadRequest, err.Error()}
	}

	wLog.UserID = userID

	if err := s.storage.CreateWorkoutLog(wLog); err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	return writeJSON(w, http.StatusCreated, nil)
}

func (s *server) handleGetWorkoutLogs(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(keyUserID).(uint)

	wLogs, err := s.storage.GetWorkoutLogs(userID)
	if err != nil {
		return apiError{http.StatusInternalServerError, err.Error()}
	}

	return writeJSON(w, http.StatusFound, wLogs)
}
