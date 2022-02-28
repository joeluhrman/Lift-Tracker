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

	err := backend.CreateProfile("joe")
	if err != nil {
		fmt.Println(err)
	}

	/*var testEntry backend.WorkoutEntry
	testEntry.Date = "12/5/3033"
	testEntry.Notes = "oh yeah now we're talkin pal"
	testEntry.Exercises = make([]backend.ExerciseEntry, 3)
	testEntry.Exercises[0].Name = "Squat"
	testEntry.Exercises[1].Name = "Bench"
	testEntry.Exercises[2].Name = "Deadlift"
	testEntry.Exercises[0].Sets = make([]backend.SetGrp, 1)
	testEntry.Exercises[1].Sets = make([]backend.SetGrp, 1)
	testEntry.Exercises[2].Sets = make([]backend.SetGrp, 2)
	testEntry.Exercises[0].Sets[0] = backend.SetGrp{Number: 3, Reps: 5, Weight: 315}
	testEntry.Exercises[1].Sets[0] = backend.SetGrp{Number: 3, Reps: 5, Weight: 225}
	testEntry.Exercises[2].Sets[1] = backend.SetGrp{Number: 1, Reps: 5, Weight: 355}
	testEntry.Exercises[2].Sets[0] = backend.SetGrp{Number: 1, Reps: 3, Weight: 385}
	testEntry.Exercises[0].Notes = "nice"
	testEntry.Exercises[1].Notes = "good"
	testEntry.Exercises[2].Notes = "two words"*/

	var testTemplate backend.WorkoutTemplate
	testTemplate.Name = "Max day"
	testTemplate.Exercises = make([]backend.ExerciseEntry, 3)
	testTemplate.Exercises[0].Name = "Bench"
	testTemplate.Exercises[1].Name = "Chins"
	testTemplate.Exercises[2].Name = "Rows"
	testTemplate.Exercises[0].Sets = make([]backend.SetGrp, 3)
	testTemplate.Exercises[0].Sets[0] = backend.SetGrp{Weight: 275, Number: 1, Reps: 1}
	testTemplate.Exercises[0].Sets[1] = backend.SetGrp{Weight: 295, Number: 1, Reps: 1}
	testTemplate.Exercises[0].Sets[1] = backend.SetGrp{Weight: 315, Number: 1, Reps: 1}
	testTemplate.Exercises[1].Sets = make([]backend.SetGrp, 1)
	testTemplate.Exercises[1].Sets[0] = backend.SetGrp{Weight: 0, Number: 3, Reps: 8}
	testTemplate.Exercises[2].Sets = make([]backend.SetGrp, 1)
	testTemplate.Exercises[2].Sets[0] = backend.SetGrp{Weight: 185, Number: 3, Reps: 5}

	err = backend.SaveWorkoutTemplate("joe", testTemplate)
	if err != nil {
		fmt.Println(err)
	}
}
