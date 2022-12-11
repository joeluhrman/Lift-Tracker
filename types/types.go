// Contains types used across multiple packages.
package types

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	IsAdmin        bool   `json:"is_admin"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username string, hashedPassword string, isAdmin bool) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		IsAdmin:        isAdmin,
	}
}

const SessionKey = "session"

type Session struct {
	UserID int
	Token  string

	CreatedAt time.Time
	UpdatedAt time.Time
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
