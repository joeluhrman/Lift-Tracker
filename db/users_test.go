package db

import "testing"

func Test_CreateUser_Valid(t *testing.T) {
	user := NewUser("jaluhrman", "123", false)
	err := CreateUser(user)

	if err != nil {
		t.Error(err)
	}
}
