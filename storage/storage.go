package storage

import (
	"os"

	"github.com/joeluhrman/Lift-Tracker/types"
)

type Storage interface {
	InsertUser(user *types.User, isAdmin bool) error
	InsertSession(s *types.Session) error
}

// not sure where to put this, can't put in main package because need for testing unfortunately
func MustReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}
