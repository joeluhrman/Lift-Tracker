package db

const (
	tableUser = "users"
)

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

func CreateUser(user *User) error {
	statement := "INSERT INTO " + tableUser + " (username, password, is_admin) VALUES ($1, $2, $3)"

	_, err := conn.Exec(statement, user.Username, user.Password, user.IsAdmin)

	return err
}
