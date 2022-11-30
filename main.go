package main

import "github.com/joeluhrman/Lift-Tracker/db"

const (
	dbDriver = "sqlite3"
	dbPath   = "database.db"
)

func main() {
	err := db.Init(&db.Config{
		Driver: dbDriver,
		Path:   dbPath,
	})
	if err != nil {
		panic(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()
}
