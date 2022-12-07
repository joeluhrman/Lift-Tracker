package db

import (
	"os"
	"testing"
)

var (
	testDBApiKey = string(MustReadFile("./api_keys/api_key_test.txt"))
	testDBConfig = &Config{
		Driver: "pgx",
		Path:   "postgresql://jaluhrman:" + testDBApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test",
	}

	tables = []string{tableUser}
)

func clearAllTables() {
	for _, table := range tables {
		clearTable(table)
	}
}

func clearTable(tName string) {
	_, err := conn.Exec("DELETE FROM " + tName)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	MustConnect(testDBConfig)
	clearAllTables()
	code := m.Run()
	os.Exit(code)
}

func Test_InsertUser_Valid(t *testing.T) {
	user := NewUser("jaluhrman", "123", false)

	err := InsertUser(user)
	if err != nil {
		t.Error(err)
	}

	clearTable(tableUser)
}
