package model

import (
	"database/sql"
	"testing"
)

// Test when all conditions are valid.
func Test_CreateUser(t *testing.T) {
	defer cleanUpDB()

	const USERNAME = "jaluhrman"
	const PASSWORD = "random_password"

	user := &User{
		Username: sql.NullString{
			String: USERNAME,
			Valid:  true,
		},
	}

	userPassword := &UserPassword{
		Password: sql.NullString{
			String: PASSWORD,
			Valid:  true,
		},
	}

	err := CreateUser(user, userPassword)
	if err != nil {
		t.Error(err)
	}

	if user.ID != userPassword.ID {
		t.Errorf("user and userPassword ids are not the same")
	}
}

// Test err is returned when username is nil.
func Test_CreateUser_NoUsername(t *testing.T) {
	defer cleanUpDB()

	const PASSWORD = "random_password"

	userPassword := &UserPassword{
		Password: sql.NullString{
			String: PASSWORD,
			Valid:  true,
		},
	}

	err := CreateUser(&User{}, userPassword)
	if err == nil {
		t.Errorf("error should have been returned when username is nil")
	}
}

// Test err is returned when password is nil.
func Test_CreateUser_NoPassword(t *testing.T) {
	defer cleanUpDB()

	const USERNAME = "jaluhrman"

	user := &User{
		Username: sql.NullString{
			String: USERNAME,
			Valid:  true,
		},
	}

	err := CreateUser(user, &UserPassword{})
	if err == nil {
		t.Errorf("error should have been returned when password is nil")
	}
}

// Test that UserPassword's ID is set to User's ID correctly before
// being created.
func Test_CreateUser_PasswordIDAlreadySet(t *testing.T) {
	defer cleanUpDB()

	const USERNAME = "jaluhrman"
	const PASSWORD = "random_password"

	user := &User{
		Username: sql.NullString{
			String: USERNAME,
			Valid:  true,
		},
	}

	userPassword := &UserPassword{
		Model: Model{
			ID: 5,
		},
		Password: sql.NullString{
			String: PASSWORD,
			Valid:  true,
		},
	}

	err := CreateUser(user, userPassword)
	if err != nil {
		t.Error(err)
	}

	if user.ID != userPassword.ID {
		t.Errorf("user and userPassword ids are not the same")
	}
}

// Test that err is returned if User with the same username already exists.
func Test_CreateUser_UsernameConflict(t *testing.T) {
	defer cleanUpDB()

	const USERNAME = "jaluhrman"
	const PASSWORD = "random_password"

	user := &User{
		Username: sql.NullString{
			String: USERNAME,
			Valid:  true,
		},
	}
	DBConn.Create(user)

	userPassword := &UserPassword{
		Password: sql.NullString{
			String: PASSWORD,
			Valid:  true,
		},
	}

	err := CreateUser(user, userPassword)
	if err == nil {
		t.Errorf("error should have been returned when username is already taken")
	}
}

// Test that err is returned and User is deleted from DB if a password with the User's id
// already exists in db.
func Test_CreateUser_UserPasswordIDConflict(t *testing.T) {
	defer cleanUpDB()

	const ID = 3
	const USERNAME = "jaluhrman"
	const PASSWORD = "random_password"

	user := &User{
		Model: Model{
			ID: ID,
		},
		Username: sql.NullString{
			String: USERNAME,
			Valid:  true,
		},
	}

	userPassword := &UserPassword{
		Model: Model{
			ID: user.ID,
		},
		Password: sql.NullString{
			String: PASSWORD,
			Valid:  true,
		},
	}
	DBConn.Create(userPassword)

	err := CreateUser(user, userPassword)
	if err == nil {
		t.Errorf("error should have been returned when password with User's ID already exists")
	}
}
