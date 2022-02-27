package main

import (
	"fmt"

	"github.com/joeluhrman/Lift-Tracker/backend"
)

func main() {
	// just testing stuff for now

	// testing static html files
	/*port := os.Getenv("PORT")
	fs := http.FileServer(http.Dir("./web/templates"))
	http.Handle("/", fs)

	log.Println("Listening on :" + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}*/

	err := backend.CreateProfile("bob")
	if err != nil {
		fmt.Println(err)
	}

	/*var testEntry backend.WorkoutEntry
	testEntry.Date = "12/5/3033"
	testEntry.Notes = "oh yeah now we're talkin pal"

	err := backend.SaveWorkoutEntry("joe", testEntry)
	if err != nil {
		fmt.Println(err)
	}*/
}
