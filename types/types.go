// Contains types used across multiple packages.
package types

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"` // not stored in db, just used for login
	HashedPassword string `json:"hashed_password"`
	IsAdmin        bool   `json:"is_admin"`

	Metadata
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

	Metadata
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

type Workout struct {
	ID        int
	UserID    int
	Name      string
	Exercises []Exercise
	Notes     string

	Metadata
}

type Exercise struct {
	ID        int
	WorkoutID int
	Name      string
	SetGroups []SetGroup
	Notes     string

	Metadata
}

type SetGroup struct {
	ID         int
	ExerciseID int

	Sets   int
	Reps   int
	Weight int

	Metadata
}
