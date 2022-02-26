package backend

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
