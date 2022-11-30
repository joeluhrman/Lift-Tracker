package main

import "github.com/joeluhrman/Lift-Tracker/db"

const (
	dbDriver = "sqlite3"
	dbPath   = "./database.db"
)

func main() {
	db.MustInit(&db.Config{
		Driver: dbDriver,
		Path:   dbPath,
	})

	defer db.MustClose()
}
