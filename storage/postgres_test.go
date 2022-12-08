package storage

import (
	"os"
	"testing"

	"github.com/joeluhrman/Lift-Tracker/types"
)

var (
	testPGDriver  = "pgx"
	testPGApiKey  = string(MustReadFile("../api_keys/api_key_test.txt"))
	testPGURL     = "postgresql://jaluhrman:" + testPGApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test"
	testPGStorage = NewPostgresStorage(testPGDriver, testPGURL)

	tables = []string{tableUser}
)

func (p *PostgresStorage) clearAllTables() {
	for _, table := range tables {
		p.clearTable(table)
	}
}

func (p *PostgresStorage) clearTable(tName string) {
	_, err := p.conn.Exec("DELETE FROM " + tName)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	testPGStorage.MustConnect()
	testPGStorage.clearAllTables()
	code := m.Run()
	os.Exit(code)
}

func Test_InsertUser_Valid(t *testing.T) {
	user := types.NewUser("jaluhrman", "123", false)

	err := testPGStorage.InsertUser(user, false)
	if err != nil {
		t.Error(err)
	}

	testPGStorage.clearTable(tableUser)
}
