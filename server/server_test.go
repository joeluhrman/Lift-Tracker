package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/joeluhrman/Lift-Tracker/types"
)

const (
	wrongCodef = "code was %d, should have been %d"
)

type testStorage struct{}

func (t *testStorage) InsertUser(user *types.User, isAdmin bool) error {
	return nil
}

var (
	ts *testStorage
)

func newTestServer() *Server {
	s := New("", ts, nil)
	s.router = chi.NewRouter()
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
	s := newTestServer()
	method := http.MethodPost
	endpoint := routeApiV1 + endCreateAcc
	successCode := http.StatusAccepted
	badJSONCode := http.StatusBadRequest
	badPasswordCode := http.StatusNotAcceptable

	// Bad JSON
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, s.router)
		if rec.Code != badJSONCode {
			t.Errorf(wrongCodef, rec.Code, badJSONCode)
		}
	}()

	// Password doesn't meet requirements
	func() {
		user := types.NewUser("jaluhrman", "", false)

		json, _ := json.Marshal(user)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, s.router)
		if rec.Code != badPasswordCode {
			t.Errorf(wrongCodef, rec.Code, badPasswordCode)
		}
	}()

	// Success case
	func() {
		user := types.NewUser("jaluhrman", "123", false)

		json, _ := json.Marshal(user)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, s.router)
		if rec.Code != successCode {
			t.Errorf(wrongCodef, rec.Code, successCode)
		}
	}()
}
