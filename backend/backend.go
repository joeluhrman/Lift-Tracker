// handles internal stuff and writing to files/storing data

package backend

import (
	"fmt"
	"os"
)

type MuscleGroup string

const (
	Legs MuscleGroup = "Legs"
	Push MuscleGroup = "Push"
	Pull MuscleGroup = "Pull"
)

type ExerciseType struct {
	Name   string
	MscGrp MuscleGroup
}

type WorkoutTemplate struct {
	Exercises []ExerciseEntry
}

type ExerciseEntry struct {
	Type  ExerciseType
	Sets  int
	Reps  int
	Notes string
}

type WorkoutEntry struct {
	Date      string
	Exercises []ExerciseEntry
	Notes     string
}

type Profile struct {
	Username string
}

func CreateProfile(username string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.MkdirAll(path+"/profiles/"+username, 0755)
	return err
}

func SaveWorkoutEntry(username string, workout WorkoutEntry) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	filepath := path + "/profiles/" + username + "/workouthistory.txt"
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	text := workout.Date + ";"
	for i := 0; i < len(workout.Exercises); i++ {
		name := workout.Exercises[i].Type.Name
		sets := workout.Exercises[i].Sets
		reps := workout.Exercises[i].Reps
		notes := workout.Exercises[i].Notes
		text += name + "|" + fmt.Sprint(sets) + "|" + fmt.Sprint(reps) + "|" + notes + ";"
	}
	text += workout.Notes

	_, err = f.WriteString(text + "\n")
	return err
}
