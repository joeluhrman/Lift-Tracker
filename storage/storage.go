// Contains Storage interface with methods for database access and the Postgres implementation.
package storage

import (
	"github.com/joeluhrman/Lift-Tracker/types"
)

type Storage interface {
	InsertUser(user *types.User, isAdmin bool) error
	InsertSession(s *types.Session) error
	AuthenticateUser(username string, password string) (int, error)
}
