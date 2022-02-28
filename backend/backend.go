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
	Weight int
	Number int
	Reps   int
}

type ExerciseType struct {
	Name   string
	MscGrp MuscleGroup
}

type ExerciseEntry struct {
	Name  string
	Sets  []SetGrp
	Notes string
}

type WorkoutTemplate struct {
	Name      string
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

func ExercisesToFileFormat(exercises []ExerciseEntry) string {
	var text string
	for i := 0; i < len(exercises); i++ {
		name := exercises[i].Name
		text += ";" + name

		for j := 0; j < len(exercises[i].Sets); j++ {
			number := exercises[i].Sets[j].Number
			weight := exercises[i].Sets[j].Weight
			reps := exercises[i].Sets[j].Reps
			divider := "|"
			if j > 0 {
				divider = "/"
			}
			text += divider + fmt.Sprint(weight) + "," + fmt.Sprint(number) + "," + fmt.Sprint(reps)
		}

		if exercises[i].Notes != "" {
			text += "|" + exercises[i].Notes
		}

	}

	return text
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

	text := workout.Date + ExercisesToFileFormat(workout.Exercises) + ";" + workout.Notes

	_, err = f.WriteString(text + "\n")
	return err
}

func SaveWorkoutTemplate(username string, workout WorkoutTemplate) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	filepath := path + "/profiles/" + username + "/workouttemplates.txt"
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	text := workout.Name + ";" + ExercisesToFileFormat(workout.Exercises)

	_, err = f.WriteString(text + "\n")
	return err
}
