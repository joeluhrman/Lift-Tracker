package types

type User struct {
	Id       int    `json:"id"`
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
