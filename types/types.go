// Contains types used across multiple packages.
package types

import (
	"image"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type metadata struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type model struct {
	metadata
	ID uint `json:"id"`
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"` // not stored in db, just used for login
	HashedPassword string `json:"hashed_password"`
	IsAdmin        bool   `json:"is_admin"`

	metadata
}

func NewUser(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

const SessionKey = "session"

type Session struct {
	UserID int
	Token  string

	metadata
}

func NewSession(userID int) *Session {
	token := uuid.New().String()

	return &Session{
		Token:  token,
		UserID: userID,
	}
}

func (s *Session) Cookie() *http.Cookie {
	return &http.Cookie{
		Name:  SessionKey,
		Value: s.Token,
	}
}

type SetGroup struct {
	model
	ExerciseID uint
	Sets       uint
	Reps       uint
	Weight     uint
}

type PPLType string

const (
	Push PPLType = "push"
	Pull         = "pull"
	Legs         = "legs"
)

type MuscleGroup string

// incomplete
const (
	Calves     MuscleGroup = "calves"
	Hamstrings             = "hamstrings"
	Quads                  = "quads"
	Adductors              = "adductors"
	Abductors              = "abductors"
	Core                   = "core"
	LowBack                = "lower back"
)

// either a default or custom exercise type (not a logged exercise)
type ExerciseType struct {
	ID          uint        `json:"id"`
	UserID      uint        `json:"user_id"`
	IsDefault   bool        `json:"is_default"`
	Name        string      `json:"name"`
	Image       image.Image `json:"image"` // png only for now
	PPLType     PPLType     `json:"ppl_type"`
	MuscleGroup MuscleGroup `json:"muscle_group"`

	metadata
}

func NewExerciseType(userID uint, isDefault bool, name string, image image.Image, pplType PPLType, mscGrp MuscleGroup) *ExerciseType {
	return &ExerciseType{
		UserID:      userID,
		IsDefault:   isDefault,
		Name:        name,
		Image:       image,
		PPLType:     pplType,
		MuscleGroup: mscGrp,
	}
}

type Exercise struct {
	model
	ExerciseTypeID uint
	WorkoutID      uint
	SetGroups      []SetGroup
	Notes          string
}

type Workout struct {
	model
	UserID    uint
	Name      string
	Date      time.Time
	Exercises []Exercise
	Notes     string
}

/* commented out for now

type Setgroup struct {
	ID         int
	ExerciseID int

	Weight int
	Sets   int
	Reps   int

	metadata
}

func NewSetgroup(exerciseID, weight, sets, reps int) *Setgroup {
	return &Setgroup{
		ExerciseID: exerciseID,
		Weight:     weight,
		Sets:       sets,
		Reps:       reps,
	}
}

type pplType string

const (
	pplPush pplType = "PUSH"
	pplPull pplType = "PULL"
	pplLegs pplType = "LEGS"
)

var (
	pplTypes = []pplType{pplPush, pplPull, pplLegs}
)

func IsPPLType(s string) bool {
	for _, v := range pplTypes {
		if s == string(v) {
			return true
		}
	}

	return false
}

type Exercise struct {
	ID        int
	WorkoutID int

	Name      string
	PPLTypes  []pplType
	Setgroups []*Setgroup
	Notes     string

	metadata
}

func NewExercise(workoutID int, name string, setgroups []*Setgroup, notes string) *Exercise {
	return &Exercise{
		WorkoutID: workoutID,
		Name:      name,
		Setgroups: setgroups,
		Notes:     notes,
	}
}

type Workout struct {
	ID     int
	UserID int

	Name      string
	Time      time.Time
	Exercises []*Exercise
	Notes     string

	metadata
}

func NewWorkout(userID int, name string, time time.Time, exercises []*Exercise, notes string) *Workout {
	return &Workout{
		UserID:    userID,
		Name:      name,
		Time:      time,
		Exercises: exercises,
		Notes:     notes,
	}
}
*/
