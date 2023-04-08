// Contains functionality for creating/starting/shutting down an HTTP server and handling HTTP requests.
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
	// routes and endpoints
	routeApiV1         = "/api/v1"
	endUser            = "/user"
	endLogin           = "/login"
	endLogout          = "/logout"
	endExerciseType    = "/exercise-type"
	endWorkoutTemplate = "/workout-template"
	endWorkoutLog      = "/workout-log"

	// key to get logged in user id from context
	keyUserID = "user_id"

	// key to get session token from cookies
	keySession = types.SessionKey
)

// A Server embeds an http.Server and has a storage.Storage for DB CRUDs.
type Server struct {
	http.Server
	storage storage.Storage
}

// New returns a *server.Server set to listen on the specified port, use the
// specified storage.Storage, and use any number of global middlewares.
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

// MustStart logs the Server's port and calls Server.ListenAndServe().
func (s *Server) MustStart() {
	log.Println("server running on port " + s.Addr)
	s.ListenAndServe()
}

// MustShutdown calls Server.Shutdown() and logs its success,
// and panics if an error occurs.
// It takes a context.Context for timeouts.
func (s *Server) MustShutdown(shutdownCtx context.Context) {
	err := s.Shutdown(shutdownCtx)
	if err != nil {
		panic(err)
	}

	log.Println("server successfully shutdown")
}

// setupEndpoints binds a *chi.Mux to the Server's endpoints.
func (s *Server) setupEndpoints(router *chi.Mux) {
	router.Route(routeApiV1, func(r chi.Router) {
		r.Post(endUser, s.handleCreateUser)
		r.Post(endLogin, s.handleLogin)

		// endpoints requiring session authentication
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

// writeJSON writes a status code and JSON data to an http.Responsewriter.
// If an encoding error occurs, it rewrites the status with http.StatusInternalServerError.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// setSession sets a *types.Session in the cookies of an http.ResponseWriter.
// Only the token is stored in cookies.
func setSession(s *types.Session, w http.ResponseWriter) {
	http.SetCookie(w, s.Cookie())
}

// getSessionToken gets the token for the current session from
// an *http.Request.
// Returns an error if the session cookie is not found.
func getSessionToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(types.SessionKey)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

// middlewareAuthSession is a simple middleware for authenticating the
// current session.
//
// If session is valid, it embeds the user id for the session into the
// *http.Request's context and serves the next handler.
//
// It writes status code http.StatusUnauthorized and the error if the
// session cookie is not found or if the session cannot be authenticated in storage.
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

/*
// handleIsLoggedIn responds with http.StatusOK and the data of the
// currently logged in user. It is meant to be used in conjunction
// with middlewareAuthSession.
//
// It responds with http.StatusInternalServerError and the error
// message if for some reason the user id in context cannot be
// found in storage.
func (s *Server) handleIsLoggedIn(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)
	user, err := s.storage.GetUser(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, user)
}
*/

// A credentials is used in the creating a user and login process to
// prevent the client from sending more data then they should be able to,
// which could result in unintended side-effects.
type credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// handleCreateUser receives a credentials json and creates the
// corresponding types.User in storage.
//
// It responds with http.StatusAccepted if successful.
//
// If the json cannot be decoded it responds with http.StatusBadRequest
// and the error.
//
// If the username or password do not meet requirements, it responds
// with http.StatusNotAcceptable and the applicable error message.
//
// If there is an error creating the user in storage (assumed to be a conflict),
// it responds with http.StatusConflict and the error message.
func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
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

// handleGetUser gets the user id of the current session from
// the *http.Request's context, retrieves that user from storage,
// and responds with http.StatusFound and the types.User.
//
// If the user canot be found, responds with http.StatusInternalServerError
// and the error message.
func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)
	user, err := s.storage.GetUser(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// handleLogin receives a credentials json and creates a session that
// is saved in storage and embedded in the the response's cookies.
//
// It responds with http.StatusOK on success.
//
// If the credentials json cannot be decoded, it responds with
// http.StatusBadRequest and the error message.
//
// If the client's credentials cannot be authenticated in storage,
// it responds with http.StatusUnauthorized and the error message.
//
// If the client's old session (if it exists) cannot be deleted from
// storage or the new session cannot be created in storage, it responds
// with http.StatusInternalServerError and the error message.
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
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

// handleLogout deletes the current client's session from storage and
// clears it from their cookies.
//
// It responds with http.StatusOK on success.
//
// It responds with http.StatusNotFound and the error message if the
// session token is not found in the client's cookies.
//
// If the session cannot be deleted from storage, it responds
// with http.StatusInternalServerError and the error message.
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

// hanleGetExerciseTypes retrieves and responds with all
// types.ExerciseType from storage.
//
// Responds with http.StatusFound and the []types.ExerciseType
// json encoding on success.
//
// Responds with http.StatusInternalServerError and the error
// message if the []types.exerciseType cannot be retrieved from
// storage.
func (s *Server) handleGetExerciseTypes(w http.ResponseWriter, r *http.Request) {
	eTypes, err := s.storage.GetExerciseTypes()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusFound, eTypes)
}

// handleCreateWorkoutTemplate receives a *types.WorkoutTemplate json
// and saves it in storage with the logged in user's id.
//
// It responds with http.StatusCreated on success.
//
// It responds with http.StatusBadRequest and the error message if the
// json cannot be decoded.
//
// It responds with http.StatusInternalServerError and the error message
// if the template cannot be created.
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

	writeJSON(w, http.StatusCreated, nil)
}

// handleGetWorkoutTemplates retrieves all the types.WorkoutTemplate for the
// currently logged in user.
//
// It responds with http.StatusFound and a []types.WorkoutTemplate JSON
// if successful.
//
// If there is a storage error, it responds with http.StatusInternalServerError
// and the error message JSON.
func (s *Server) handleGetWorkoutTemplates(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)

	wTemps, err := s.storage.GetWorkoutTemplates(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusFound, wTemps)
}

// handleCreateWorkoutLog decodes from JSON and saves a
// *types.WorkoutLog in storage with the logged in user's id.
//
// If successful, it responds with http.StatusCreated.
//
// If there is a JSON decoding error, responds with http.StatusBadRequest
// and the error message JSON.
//
// If the log cannot be created in storage, it responds with
// http.StatusInternalServerError and the error message JSON.
func (s *Server) handleCreateWorkoutLog(w http.ResponseWriter, r *http.Request) {
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

// handleGetWorkoutLogs retrieves all types.WorkoutLog for the
// currently logged in user.
//
// On success, it responds with http.StatusFound and a []types.WorkoutLog
// JSON.
//
// On storage error, it responds with http.StatusInternalServerError and
// the error message JSON.
func (s *Server) handleGetWorkoutLogs(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(keyUserID).(uint)

	wLogs, err := s.storage.GetWorkoutLogs(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusFound, wLogs)
}
