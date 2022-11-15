package model

import (
	"os"
	"testing"
)

const (
	TEST_DB_PATH = "file::memory:?cache=shared"
)

// Initializes in-memory db for testing.
func initTestDB() {
	err := InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}
}

// Deletes all rows from all tables.
func cleanUpDB() {
	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM user_passwords")
}

func TestMain(m *testing.M) {
	initTestDB()
	rc := m.Run()
	os.Exit(rc)
}
