package types

import (
	"image"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
	IsAdmin        bool   `json:"isAdmin"`
	Metadata
}

const SessionKey = "session"

type Session struct {
	UserID uint
	Token  string
	Metadata
}

func NewSession(userID uint) *Session {
	token := uuid.New().String()

	return &Session{
		Token:  token,
		UserID: userID,
	}
}

func (s *Session) Cookie() *http.Cookie {
	return &http.Cookie{
		Name:  SessionKey,
		Value: s.Token,
	}
}

type PPLType string

const (
	Push PPLType = "push"
	Pull PPLType = "pull"
	Legs PPLType = "legs"
)

type MuscleGroup string

const (
	Calves     MuscleGroup = "calves"
	Hamstrings MuscleGroup = "hamstrings"
	Quads      MuscleGroup = "quads"
	Adductors  MuscleGroup = "adductors"
	Abductors  MuscleGroup = "abductors"
	Core       MuscleGroup = "core"
	LowBack    MuscleGroup = "lower back"
	Chest      MuscleGroup = "chest"
	Lats       MuscleGroup = "lats"
)

type ExerciseType struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Image       image.Image `json:"image"`
	PPLType     PPLType     `json:"pplType"`
	MuscleGroup MuscleGroup `json:"muscleGroup"`
	Metadata
}

func NewExerciseType(name string, image image.Image, pplType PPLType, mscGrp MuscleGroup) *ExerciseType {
	return &ExerciseType{
		Name:        name,
		Image:       image,
		PPLType:     pplType,
		MuscleGroup: mscGrp,
	}
}

type SetGroupLog struct {
	ID            uint    `json:"id"`
	ExerciseLogID uint    `json:"exerciseLogID"`
	Sets          uint    `json:"sets"`
	Reps          uint    `json:"reps"`
	Weight        float32 `json:"weight"`
	Metadata
}

type ExerciseLog struct {
	ID             uint          `json:"id"`
	WorkoutLogID   uint          `json:"workoutLogID"`
	ExerciseTypeID uint          `json:"exerciseTypeID"`
	Notes          string        `json:"notes"`
	SetGroupLogs   []SetGroupLog `json:"setgroupLogs"`
	Metadata
}

type WorkoutLog struct {
	ID           uint          `json:"id"`
	UserID       uint          `json:"userID"`
	Date         time.Time     `json:"date"`
	Name         string        `json:"name"`
	Notes        string        `json:"notes"`
	ExerciseLogs []ExerciseLog `json:"exerciseLogs"`
	Metadata
}

type SetGroupTemplate struct {
	ID                 uint `json:"id"`
	ExerciseTemplateID uint `json:"exerciseTemplateID"`
	Sets               uint `json:"sets"`
	Reps               uint `json:"reps"`
	Metadata
}

type ExerciseTemplate struct {
	ID                uint               `json:"id"`
	WorkoutTemplateID uint               `json:"workoutTemplateID"`
	ExerciseTypeID    uint               `json:"exerciseTypeID"`
	SetGroupTemplates []SetGroupTemplate `json:"setgroupTemplates"`
	Metadata
}

type WorkoutTemplate struct {
	ID                uint               `json:"id"`
	UserID            uint               `json:"userID"`
	Name              string             `json:"name"`
	ExerciseTemplates []ExerciseTemplate `json:"exerciseTemplates"`
	Metadata
}
