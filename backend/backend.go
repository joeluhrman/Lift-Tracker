// handles internal stuff and writing to files/storing data

package backend

import (
	"fmt"
	"os"
)

type MuscleGroup string

const (
	LEGS MuscleGroup = "Legs"
	PUSH MuscleGroup = "Push"
	PULL MuscleGroup = "Pull"

	EXERCISETEMPLATES string = "exercisetemplates.txt"
	WORKHISTORY       string = "workouthistory.txt"
	WORKTEMPLATES     string = "workouttemplates.txt"
)

type SetGrp struct {
	Weight  int
	NumSets int
	Reps    int
}

type Exercise struct {
	Name   string
	MscGrp MuscleGroup
	Sets   []SetGrp
	Notes  string
}

type Workout struct {
	Name      string
	Date      string
	Exercises []Exercise
	Notes     string
}

type Profile struct {
	Username string
	Password string
}

func CreateProfile(username string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.MkdirAll(path+"/profiles/"+username, 0755)
	return err
}

func exercisesToFileFormat(exercises []Exercise, isTemplate bool) string {
	var text string
	text += ";"
	for i := 0; i < len(exercises); i++ {
		name := exercises[i].Name
		mscgrp := exercises[i].MscGrp

		if i > 0 {
			text += "%"
		}

		text += name + "|" + string(mscgrp) + "|"

		for j := 0; j < len(exercises[i].Sets); j++ {
			number := exercises[i].Sets[j].NumSets
			weight := exercises[i].Sets[j].Weight
			reps := exercises[i].Sets[j].Reps
			if j > 0 {
				text += "/"
			}
			text += fmt.Sprint(weight) + "," + fmt.Sprint(number) + "," + fmt.Sprint(reps)
		}

		if !isTemplate && exercises[i].Notes != "" {
			text += "|" + exercises[i].Notes
		}

	}

	return text
}

func SaveExercise(username string, exercise *Exercise) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	filepath := path + "/profiles/" + username + "/" + EXERCISETEMPLATES
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	text := exercise.Name + "|" + string(exercise.MscGrp)

	_, err = f.WriteString(text + "\n")
	return err
}

// saves a workout to a text file and omits some data if saving a template
func SaveWorkout(username string, workout *Workout, isTemplate bool) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	// determine whether we're saving an entry or a template
	location := WORKHISTORY
	if isTemplate {
		location = WORKTEMPLATES
	}

	// open appropriate file
	filepath := path + "/profiles/" + username + "/" + location
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	// concatenate data to print to file depending on whehter entry or template
	var text string
	text += workout.Name
	if location == WORKHISTORY {
		text += ";" + workout.Date
	}
	text += exercisesToFileFormat(workout.Exercises, isTemplate)
	if location == WORKHISTORY {
		text += ";" + workout.Notes
	}

	_, err = f.WriteString(text + "\n")
	return err
}
