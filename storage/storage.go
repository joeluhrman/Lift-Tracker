// Contains Storage interface with methods for database access and the Postgres implementation.
package storage

import (
	"github.com/joeluhrman/Lift-Tracker/types"
)

type Storage interface {
	CreateUser(user *types.User, isAdmin bool) error
	AuthenticateUser(username string, password string) (int, error)

	CreateSession(s *types.Session) error
	DeleteSessionByUserID(userID int) error
	DeleteSessionByToken(token string) error
	AuthenticateSession(token string) (int, error)

	CreateExerciseType(exerciseType *types.ExerciseType) error
	GetExerciseTypes(userID uint) ([]types.ExerciseType, error)
}
