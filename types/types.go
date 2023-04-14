// Contains types used across multiple packages.
package types

import (
	"image"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// A types.Metadata contains metadata for a type/record
// in storage.
type Metadata struct {
	// date/time record was created
	CreatedAt time.Time `json:"createdAt"`

	// date/time record was updated
	UpdatedAt time.Time `json:"updatedAt"`
}

// A types.User represents a user's account information
// in storage.
type User struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
	IsAdmin        bool   `json:"isAdmin"`

	Metadata
}

// key for a session token in cookies
const SessionKey = "session"

// A types.Session is used in storage and cookies
// to authenticate the currently logged in user.
type Session struct {
	UserID uint   // id of the user
	Token  string // unique UUID for the session

	Metadata
}

// types.NewSession returns a *types.Session with
// the passed user id and a unique UUID token.
func NewSession(userID uint) *Session {
	token := uuid.New().String()

	return &Session{
		Token:  token,
		UserID: userID,
	}
}

// types.Session.Cookie returns an *http.Cookie
// representation of the session key and token.
func (s *Session) Cookie() *http.Cookie {
	return &http.Cookie{
		Name:  SessionKey,
		Value: s.Token,
	}
}

// A PPLType represents the Push/Pull/Legs paradigm
// for lifting programming.
type PPLType string

// PPLType
const (
	Push PPLType = "push"
	Pull PPLType = "pull"
	Legs PPLType = "legs"
)

// A MuscleGroup represents a specific major human
// muscle group that an exercise targets.
type MuscleGroup string

// MuscleGroup
//
// INCOMPLETE
const (
	Calves     MuscleGroup = "calves"
	Hamstrings MuscleGroup = "hamstrings"
	Quads      MuscleGroup = "quads"
	Adductors  MuscleGroup = "adductors"
	Abductors  MuscleGroup = "abductors"
	Core       MuscleGroup = "core"
	LowBack    MuscleGroup = "lower back"
)

// A types.ExerciseType represents a specific exercise
// (i.e. a squat, benchpress, deadlift, etc.) in storage.
//
// It is used to ensure that all template and logged
// exercises created by users are mapped to a valid
// exercise in storage. This allows accurate/consistent
// analytics on user workout data.
//
// Currently, only default exercise types are supported,
// but in the future it is intended for users to be able
// to create their own custom exercise types (provided they
// do not conflict with a default one.)
type ExerciseType struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Image       image.Image `json:"image";"omitempty"` // png only for now
	PPLType     PPLType     `json:"pplType"`
	MuscleGroup MuscleGroup `json:"muscleGroup"`

	Metadata
}

// types.NewExerciseType returns a new *types.ExerciseType.
func NewExerciseType(name string, image image.Image, pplType PPLType, mscGrp MuscleGroup) *ExerciseType {
	return &ExerciseType{
		Name:        name,
		Image:       image,
		PPLType:     pplType,
		MuscleGroup: mscGrp,
	}
}

// A types.SetGroupLog represents a logged setgroup part
// of a greater logged exercise and workout in storage.
//
// It allows a user to log an exercise with different
// groupings of sets instead of one set/rep/weight
// scheme. For example, if a user wanted to squat
// 3x5 w/ 315lbs then 2x8 w/ 225lbs, they could
// log one squat exercise with two setgroups instead
// of having to log two squat exercises for each
// set/rep/weight scheme.
type SetGroupLog struct {
	ID            uint    `json:"id"`
	ExerciseLogID uint    `json:"exerciseLogID"`
	Sets          uint    `json:"sets"`
	Reps          uint    `json:"reps"`
	Weight        float32 `json:"weight"`

	Metadata
}

// A *types.ExerciseLog represents a logged exercise
// in storage.
type ExerciseLog struct {
	ID           uint `json:"id"`
	WorkoutLogID uint `json:"workoutLogID"`

	// ID of the exercise type it corresponds to
	ExerciseTypeID uint `json:"exerciseTypeID"`

	// notes on how the exercise went etc.
	Notes string `json:"notes"`

	// can have multiple setgroups per exercise
	SetGroupLogs []SetGroupLog `json:"setgroupLogs"`

	Metadata
}

// A *types.WorkoutLog represents a logged workout in
// storage.
type WorkoutLog struct {
	ID     uint      `json:"id"`
	UserID uint      `json:"userID"`
	Date   time.Time `json:"date"`
	Name   string    `json:"name"`

	// notes on how the workout in general went
	Notes        string        `json:"notes"`
	ExerciseLogs []ExerciseLog `json:"exerciseLogs"`

	Metadata
}

// A types.SetGroupTemplate represents a template
// of a setgroup as part of a greater exercise template
// and workout template. It provides the same benefits
// of a types.SetGroupLog, just for templates instead
// of logs (it doesn't have a weight field).
type SetGroupTemplate struct {
	ID                 uint `json:"id"`
	ExerciseTemplateID uint `json:"exerciseTemplateID"`
	Sets               uint `json:"sets"`
	Reps               uint `json:"reps"`

	Metadata
}

// A types.ExerciseTemplate represents a template
// for an exercise as part of a greater workout
// template in storage.
type ExerciseTemplate struct {
	ID                uint `json:"id"`
	WorkoutTemplateID uint `json:"workoutTemplateID"`

	// id of the exercise type it corresponds to
	ExerciseTypeID uint `json:"exerciseTypeID"`

	// can have multiple setgroups
	SetGroupTemplates []SetGroupTemplate `json:"setgroupTemplates"`

	Metadata
}

// A types.WorkoutTemplate represents a template of
// for a workout in storage.
type WorkoutTemplate struct {
	ID                uint               `json:"id"`
	UserID            uint               `json:"userID"`
	Name              string             `json:"name"`
	ExerciseTemplates []ExerciseTemplate `json:"exerciseTemplates"`

	Metadata
}
