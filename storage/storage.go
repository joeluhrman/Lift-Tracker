package storage

import (
	"github.com/joeluhrman/Lift-Tracker/types"
)

type Storage interface {
	CreateUser(user *types.User, password string) error
	AuthenticateUser(username string, password string) (uint, error)
	GetUser(userID uint) (types.User, error)

	CreateSession(s types.Session) error
	DeleteSessionByUserID(userID uint) error
	DeleteSessionByToken(token string) error
	AuthenticateSession(token string) (uint, error)

	CreateExerciseType(exerciseType *types.ExerciseType) error
	GetExerciseTypes() ([]types.ExerciseType, error)

	CreateWorkoutTemplate(workoutTemplate *types.WorkoutTemplate) error
	GetWorkoutTemplates(userID uint) ([]types.WorkoutTemplate, error)

	CreateWorkoutLog(wLog *types.WorkoutLog) error
	GetWorkoutLogs(userID uint) ([]types.WorkoutLog, error)
}
