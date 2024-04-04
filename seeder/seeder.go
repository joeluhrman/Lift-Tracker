package main

import (
	"github.com/joeluhrman/Lift-Tracker/storage"
	"github.com/joeluhrman/Lift-Tracker/types"
)

func main() {
	var (
		pgDriver = "pgx"
		pgURL    = string(storage.MustReadFile("./pg_conn_string.txt"))
	)

	pgStore := storage.NewPostgres(pgDriver, pgURL)
	pgStore.MustOpen()
	defer pgStore.MustClose()

	seedExerciseTypes(pgStore)
}

func seedExerciseTypes(store storage.Storage) {
	store.CreateExerciseType(&types.ExerciseType{
		Name:         "Barbell Squat",
		PPLTypes:     types.PPLTypesFromStrings([]string{string(types.PPLTypeLegs)}),
		MuscleGroups: types.MuscleGroupsFromStrings([]string{string(types.MscGrpQuads)}),
	})

	store.CreateExerciseType(&types.ExerciseType{
		Name:         "Barbell Bench",
		PPLTypes:     types.PPLTypesFromStrings([]string{string(types.PPLTypePush)}),
		MuscleGroups: types.MuscleGroupsFromStrings([]string{string(types.MscGrpChest)}),
	})

	store.CreateExerciseType(&types.ExerciseType{
		Name:         "Deadlift",
		PPLTypes:     types.PPLTypesFromStrings([]string{string(types.PPLTypeLegs), string(types.PPLTypePull)}),
		MuscleGroups: types.MuscleGroupsFromStrings([]string{string(types.MscGrpHamstrings)}),
	})
}
