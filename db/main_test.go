package db

import (
	"os"
	"testing"
)

func clearAllTables() {
	_, err := conn.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}

func clearTable(tName string) {
	_, err := conn.Exec("DELETE FROM " + tName)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	os.Chdir("..")
	MustConnect(TestDBConfig)
	clearAllTables()
	code := m.Run()
	os.Exit(code)
}
