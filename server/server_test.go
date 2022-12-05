package server

import (
	"os"
	"testing"

	_ "github.com/joeluhrman/Lift-Tracker/testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_MustStart_Smoke(t *testing.T) {
	MustStart(TestServerConfig)
}
