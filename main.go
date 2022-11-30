package main

import "github.com/joeluhrman/Lift-Tracker/db"

func main() {
	db.Init()

	defer db.Close()
}
