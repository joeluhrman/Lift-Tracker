package main

import (
	"os"

	"github.com/joeluhrman/Lift-Tracker/db"
)

func main() {
	const (
		dbDriver     = "pgx"
		dbApiKeyPath = "./db/api_key.txt"
	)

	var (
		dbApiKey = string(mustReadFile(dbApiKeyPath))

		dbURL = db.NewURL("Lift-Tracker", "db.bit.io" /*dbPort,*/, "jaluhrman", dbApiKey)
	)

	db.MustInit(&db.Config{
		Driver: dbDriver,
		URL:    dbURL,
	})

	defer db.MustClose()
}

func mustReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}
