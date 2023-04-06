// Contains Storage interface with methods for CRUDs and the Postgres implementation.
package storage

import (
	"github.com/joeluhrman/Lift-Tracker/types"
)

// The Storage interface abstracts out the needed functionality
// for CRUDs in the database.
type Storage interface {
	// CreateUser saves a *types.User in storage, scanning the
	// id back into the pointer, and returns an
	// error if one occurred.
	// Since a types.User only has a HashedPassword field, the
	// plaintext password must be passed in as a separate param.
	CreateUser(user *types.User, password string) error

	// AuthenticateUser takes a username and password and returns
	// the user id they correspond to from storage, and an error
	// if the user could not be authenticated.
	AuthenticateUser(username string, password string) (uint, error)

	// GetUser returns a types.User from storage based on an id
	// and an error if the user could not be found.
	GetUser(userID uint) (types.User, error)

	// CreateSession saves a *types.Session in storage and returns
	// an error if there was one.
	CreateSession(s *types.Session) error

	// DeleteSessionByUserID deletes a types.Session from storage
	// with the corresponding user id. It returns an error if one
	// occurred.
	DeleteSessionByUserID(userID uint) error

	// DeleteSessionByToken delets a types.Session from storage
	// with the corresponding token, and returns an error if
	// one occurred.
	DeleteSessionByToken(token string) error

	// AuthenticateSession takes a token and returns from storage
	// the user id of the session with that token. It returns an
	// error if the session could not be found.
	AuthenticateSession(token string) (uint, error)

	// CreateExerciseType saves a *types.ExerciseType in storage,
	// scanning the id back into the pointer, and returns an error
	// if one occurred.
	//
	// CURRENTLY ONLY USED FOR DEVELOPMENT B/C CUSTOM EXERCISE
	// TYPES HAVE NOT YET BEEN IMPLEMENTED.
	CreateExerciseType(exerciseType *types.ExerciseType) error

	// GetExerciseTypes returns a []types.ExerciseType with
	// every types.ExerciseType in storage, and an error
	// if one occurred.
	GetExerciseTypes() ([]types.ExerciseType, error)

	// CreateWorkoutTemplate saves a *types.WorkoutTemplate in storage
	// and scans the id back into the pointer, and returns an error if
	// one occurred.
	CreateWorkoutTemplate(workoutTemplate *types.WorkoutTemplate) error

	// GetWorkoutTemplates returns a []types.WorkoutTemplate of all
	// the types.WorkoutTemplate in storage with the corresponding
	// user id. It returns an error if one occurred.
	GetWorkoutTemplates(userID uint) ([]types.WorkoutTemplate, error)

	// CreateWorkoutLog saves a *types.WorkoutLog in storage and scans
	// the id back into the pointer, returning an error if one occurred.
	CreateWorkoutLog(wLog *types.WorkoutLog) error

	// GetWorkoutLogs returns a []types.WorkoutLog with all the
	// types.WorkoutLog that correspond to the user id, and an error
	// if one occurred.
	GetWorkoutLogs(userID uint) ([]types.WorkoutLog, error)
}
