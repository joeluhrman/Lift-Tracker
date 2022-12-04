package db

type User struct {
	id       int
	username string
	password string

	isAdmin bool
}

func CreateUser(user *User) {

}
