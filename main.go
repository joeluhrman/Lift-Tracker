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

	/*err := backend.CreateProfile("bob")
	if err != nil {
		fmt.Println(err)
	}*/

	var testEntry backend.WorkoutEntry
	testEntry.Date = "12/5/3033"
	testEntry.Notes = "oh yeah now we're talkin pal"
	testEntry.Exercises = make([]backend.ExerciseEntry, 3)
	testEntry.Exercises[0].Type.Name = "Squat"
	testEntry.Exercises[0].Sets = 3
	testEntry.Exercises[0].Reps = 5
	testEntry.Exercises[0].Notes = "went well"
	testEntry.Exercises[1].Type.Name = "Bench"
	testEntry.Exercises[1].Sets = 3
	testEntry.Exercises[1].Reps = 5
	testEntry.Exercises[1].Notes = "good"
	testEntry.Exercises[2].Type.Name = "Deadlift"
	testEntry.Exercises[2].Sets = 3
	testEntry.Exercises[2].Reps = 5
	testEntry.Exercises[2].Notes = "oh yeah"

	err := backend.SaveWorkoutEntry("joe", testEntry)
	if err != nil {
		fmt.Println(err)
	}
}
