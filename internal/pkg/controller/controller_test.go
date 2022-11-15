package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/model"
)

const (
	TEST_DB_PATH  = "file::memory:?cache=shared"
	CON_TYPE_JSON = "application/json"
	CON_TYPE_PF   = "application/x-www-form-urlencoded"
)

// Initializes in-memory db for testing.
func initTestDB() {
	err := model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}
}

// Deletes all rows from all tables.
func cleanUpDB() {
	model.DBConn.Exec("DELETE FROM users")
	model.DBConn.Exec("DELETE FROM user_passwords")
}

// Sets up a engine for testing purposes. You can pass any number of custom
// middlewares depending on how you want to set up your test.
func setupTestEngine(middlewares ...gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)

	engine := gin.Default()

	//testStore := cookie.NewStore([]byte("test"))
	//engine.Use(sessions.Sessions("test_session", testStore))

	for _, middleware := range middlewares {
		engine.Use(middleware)
	}

	SetupEndpoints(engine)

	return engine
}

// Sends a mock HTTP request and returns *httptest.ResponseRecorder. Panics if error encountered.
func sendMockHTTPRequest(method string, endpoint string, data *bytes.Buffer, contentType string, engine *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	if data == nil {
		data = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", contentType)

	engine.ServeHTTP(w, req)

	return w
}

func TestMain(m *testing.M) {
	initTestDB()
	rc := m.Run()
	os.Exit(rc)
}
