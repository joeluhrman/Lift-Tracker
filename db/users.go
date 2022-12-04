package db

const (
	tableUser = "users"
)

type User struct {
	id       int
	username string
	password string
	isAdmin  bool
}

func NewUser(username string, password string, isAdmin bool) *User {
	return &User{
		username: username,
		password: password,
		isAdmin:  isAdmin,
	}
}

func CreateUser(user *User) error {
	statement := "INSERT INTO " + tableUser + " (username, password, is_admin) VALUES ($1, $2, $3)"

	_, err := conn.Exec(statement, user.username, user.password, user.isAdmin)

	return err
}
