package db

import (
	"os"
	"testing"

	_ "github.com/joeluhrman/Lift-Tracker/testing"
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
	MustConnect(TestDBConfig)
	clearAllTables()
	code := m.Run()
	os.Exit(code)
}
