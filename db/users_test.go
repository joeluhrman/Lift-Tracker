package db

import "testing"

func Test_InsertUser_Valid(t *testing.T) {
	user := NewUser("jaluhrman", "123", false)

	err := InsertUser(user)
	if err != nil {
		t.Error(err)
	}

	clearTable(tableUser)
}
