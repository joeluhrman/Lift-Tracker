package server

import (
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi/middleware"
)

var (
	testServerConfig = &Config{
		Port: ":3000",
		Middlewares: []func(http.Handler) http.Handler{
			middleware.Logger,
		},
	}
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_MustStart_Smoke(t *testing.T) {
	MustStart(testServerConfig)
}
