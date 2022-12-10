// Contains types used across multiple packages.
package types

import "time"

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

type Session struct {
	UserID int
	Token  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSession(token string, userID int) *Session {
	return &Session{
		Token:  token,
		UserID: userID,
	}
}
