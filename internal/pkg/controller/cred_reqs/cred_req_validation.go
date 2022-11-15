package cred_reqs

import "errors"

var (
	ERR_USERNAME_NOT_COMPLEX = errors.New("Username did not meet the complexity requirements")
	ERR_PASSWORD_NOT_COMPLEX = errors.New("Password did not meet the complexity requirements")
)

func CheckUsernameMeetsRequirements(username string) error {
	var err error

	if username == "" {
		err = ERR_USERNAME_NOT_COMPLEX
	}

	return err
}

func CheckPasswordMeetsRequirements(password string) error {
	var err error

	if password == "" {
		err = ERR_PASSWORD_NOT_COMPLEX
	}

	return err
}
