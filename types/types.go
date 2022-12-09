package types

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func NewUser(username string, password string, isAdmin bool) *User {
	return &User{
		Username: username,
		Password: password,
		IsAdmin:  isAdmin,
	}
}

type Session struct {
	UserID int
	Token  string
}

func NewSession(token string, userID int) *Session {
	return &Session{
		Token:  token,
		UserID: userID,
	}
}
