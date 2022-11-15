package cred_reqs

import "errors"

var (
	ERR_USERNAME_NOT_COMPLEX = errors.New("Username did not meet the complexity requirements")
	ERR_PASSWORD_NOT_COMPLEX = errors.New("Password did not meet the complexity requirements")
)

// Checks that a username meets the requirements. Returns error if does not.
func CheckUsername(username string) error {
	var err error

	if username == "" {
		err = ERR_USERNAME_NOT_COMPLEX
	}

	return err
}

// Checks that a password meets the requirements. Returns error if does not.
func CheckPassword(password string) error {
	var err error

	if password == "" {
		err = ERR_PASSWORD_NOT_COMPLEX
	}

	return err
}
