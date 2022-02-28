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

type SetGrp struct {
	Number int
	Reps   int
	Weight int
}

type ExerciseType struct {
	Name   string
	MscGrp MuscleGroup
}

type ExerciseEntry struct {
	Type  ExerciseType
	Sets  []SetGrp
	Notes string
}

type WorkoutTemplate struct {
	Exercises []ExerciseEntry
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

	text := workout.Date

	for i := 0; i < len(workout.Exercises); i++ {
		name := workout.Exercises[i].Type.Name
		text += ";" + name

		for j := 0; j < len(workout.Exercises[i].Sets); j++ {
			number := workout.Exercises[i].Sets[j].Number
			weight := workout.Exercises[i].Sets[j].Weight
			reps := workout.Exercises[i].Sets[j].Reps
			divider := "|"
			if j > 0 {
				divider = "/"
			}
			text += divider + fmt.Sprint(weight) + "," + fmt.Sprint(number) + "," + fmt.Sprint(reps)
		}

		text += "|" + workout.Exercises[i].Notes
	}
	text += ";" + workout.Notes

	_, err = f.WriteString(text + "\n")
	return err
}
