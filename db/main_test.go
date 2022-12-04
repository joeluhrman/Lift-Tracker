package db

import (
	"os"
	"testing"

	"github.com/joeluhrman/Lift-Tracker/utils"
)

var (
	testDBApiKey = string(utils.MustReadFile("./api_key_test.txt"))

	TestDBConfig = &Config{
		Driver: "pgx",
		Path:   "postgresql://jaluhrman:" + testDBApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test",
	}
)

func clearTables() {
	_, err := conn.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	MustConnect(TestDBConfig)
	clearTables()
	code := m.Run()
	clearTables()
	os.Exit(code)
}
