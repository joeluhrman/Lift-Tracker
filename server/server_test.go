package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/joeluhrman/Lift-Tracker/storage"
	"github.com/joeluhrman/Lift-Tracker/types"
)

const (
	wrongCodef = "code was %d, should have been %d"
)

var (
	testServer = newTestServer(&testStorage{}, nil)
)

type testStorage struct{}

func (t *testStorage) InsertUser(user *types.User, isAdmin bool) error {
	return nil
}

func (t *testStorage) AuthenticateUser(username string, password string) (int, error) {
	return 1, nil
}

func (t *testStorage) InsertSession(s *types.Session) error {
	return nil
}

func (t *testStorage) DeleteSessionByUserID(userID int) error {
	return nil
}

func (t *testStorage) DeleteSessionByToken(token string) error {
	return nil
}

func (t *testStorage) InsertWorkout(w *types.Workout) error {
	return nil
}

func newTestServer(storage storage.Storage, middlewares []middleware) *Server {
	s := New("", storage, nil)
	s.router = chi.NewRouter()

	for _, middleware := range middlewares {
		s.router.Use(middleware)
	}

	s.setupEndpoints()

	return s
}

func sendMockHTTPRequest(method string, endpoint string, data *bytes.Buffer, router *chi.Mux) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()

	if data == nil {
		data = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		panic(err)
	}

	router.ServeHTTP(rec, req)

	return rec
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_handleCreateAccount(t *testing.T) {
	method := http.MethodPost
	endpoint := routeApiV1 + endCreateAcc
	successCode := http.StatusAccepted
	badJSONCode := http.StatusBadRequest
	badPasswordCode := http.StatusNotAcceptable

	// Bad JSON
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, testServer.router)
		if rec.Code != badJSONCode {
			t.Errorf(wrongCodef, rec.Code, badJSONCode)
		}
	}()

	// Password doesn't meet requirements
	func() {
		user := types.NewUser("jaluhrman", "")

		json, _ := json.Marshal(user)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, testServer.router)
		if rec.Code != badPasswordCode {
			t.Errorf(wrongCodef, rec.Code, badPasswordCode)
		}
	}()

	// Success case
	func() {
		user := types.NewUser("jaluhrman", "123")

		json, _ := json.Marshal(user)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, testServer.router)
		if rec.Code != successCode {
			t.Errorf(wrongCodef, rec.Code, successCode)
		}
	}()
}

func Test_handleLogin(t *testing.T) {
	method := http.MethodPost
	endpoint := routeApiV1 + endLogin
	badJSONCode := http.StatusBadRequest

	// bad json
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, testServer.router)
		if rec.Code != badJSONCode {
			t.Errorf(wrongCodef, rec.Code, badJSONCode)
		}
	}()

	// success case
	func() {
		loginInfo := &types.User{
			Username:       "jaluhrman",
			HashedPassword: "123",
		}

		json, _ := json.Marshal(loginInfo)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, testServer.router)

		// check correct response code
		if rec.Code != http.StatusOK {
			t.Errorf(wrongCodef, rec.Code, http.StatusOK)
		}

		// check session cookie has been set correctly
		sessionSet := false
		cookies := rec.Result().Cookies()
		for _, cookie := range cookies {
			if cookie.Name == types.SessionKey && cookie.Value != "" {
				sessionSet = true
			}
		}
		if !sessionSet {
			t.Error("session cookie was not set")
		}
	}()
}

func Test_handleLogout(t *testing.T) {
	const (
		method   = http.MethodPost
		endpoint = routeApiV1 + endLogout
		userID   = 1
	)

	// cookies reset correctly
	func() {
		server := newTestServer(&testStorage{}, []middleware{func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.AddCookie(types.NewSession(userID).Cookie())
				next.ServeHTTP(w, r)
			})
		}})

		rec := sendMockHTTPRequest(method, endpoint, nil, server.router)
		if rec.Code != http.StatusOK {
			t.Errorf(wrongCodef, rec.Code, http.StatusOK)
		}

		cookieReset := false
		cookies := rec.Result().Cookies()
		for _, cookie := range cookies {
			if cookie.Name == types.SessionKey && cookie.Value == "" {
				cookieReset = true
			}
		}
		if !cookieReset {
			t.Error("session cookie was not reset")
		}
	}()
}
