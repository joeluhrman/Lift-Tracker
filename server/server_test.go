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
	testServer = newTestServer(&testPGStorage{})
)

func newTestServer(storage storage.Storage) *Server {
	s := New("", storage, nil)
	s.router = chi.NewRouter()
	s.setupEndpoints()

	return s
}

type testPGStorage struct{}

func (t *testPGStorage) InsertUser(user *types.User, isAdmin bool) error {
	return nil
}

func (t *testPGStorage) InsertSession(s *types.Session) error {
	return nil
}

func (t *testPGStorage) AuthenticateUser(username string, password string) (int, error) {
	return 1, nil
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
	badPasswordCode := http.StatusNotAcceptable

	// Bad JSON
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, testServer.router)
		if rec.Code != codeErrBadJSON {
			t.Errorf(wrongCodef, rec.Code, codeErrBadJSON)
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

	// bad json
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, testServer.router)
		if rec.Code != codeErrBadJSON {
			t.Errorf(wrongCodef, rec.Code, codeErrBadJSON)
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
		if rec.Code != codeSuccLogin {
			t.Errorf(wrongCodef, rec.Code, codeSuccLogin)
		}
	}()
}
