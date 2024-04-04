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
	pplTypes, _ := types.PPLTypesFromStrings([]string{string(types.PPLTypeLegs)})
	store.CreateExerciseType(&types.ExerciseType{
		Name:         "Barbell Squat",
		PPLTypes:     pplTypes,
		MuscleGroups: types.MuscleGroupsFromStrings([]string{string(types.MscGrpQuads)}),
	})

	pplTypes, _ = types.PPLTypesFromStrings([]string{string(types.PPLTypePush)})
	store.CreateExerciseType(&types.ExerciseType{
		Name:         "Barbell Bench",
		PPLTypes:     pplTypes,
		MuscleGroups: types.MuscleGroupsFromStrings([]string{string(types.MscGrpChest)}),
	})

	pplTypes, _ = types.PPLTypesFromStrings([]string{string(types.PPLTypeLegs), string(types.PPLTypePull)})
	store.CreateExerciseType(&types.ExerciseType{
		Name:         "Deadlift",
		PPLTypes:     pplTypes,
		MuscleGroups: types.MuscleGroupsFromStrings([]string{string(types.MscGrpHamstrings)}),
	})
}
