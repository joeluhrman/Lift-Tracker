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
	ID          uint        `json:"id,string"`
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
	ID            uint    `json:"id,string"`
	ExerciseLogID uint    `json:"exerciseLogID,string"`
	Sets          uint    `json:"sets,string"`
	Reps          uint    `json:"reps,string"`
	Weight        float32 `json:"weight"`
	Metadata
}

type ExerciseLog struct {
	ID             uint          `json:"id,string"`
	WorkoutLogID   uint          `json:"workoutLogID,string"`
	ExerciseTypeID uint          `json:"exerciseTypeID,string"`
	Notes          string        `json:"notes"`
	SetGroupLogs   []SetGroupLog `json:"setgroupLogs"`
	Metadata
}

type WorkoutLog struct {
	ID           uint          `json:"id,string"`
	UserID       uint          `json:"userID,string"`
	Date         time.Time     `json:"date"`
	Name         string        `json:"name"`
	Notes        string        `json:"notes"`
	ExerciseLogs []ExerciseLog `json:"exerciseLogs"`
	Metadata
}

type SetGroupTemplate struct {
	ID                 uint `json:"id,string"`
	ExerciseTemplateID uint `json:"exerciseTemplateID,string"`
	Sets               uint `json:"sets,string"`
	Reps               uint `json:"reps,string"`
	Metadata
}

type ExerciseTemplate struct {
	ID                uint               `json:"id,string"`
	WorkoutTemplateID uint               `json:"workoutTemplateID,string"`
	ExerciseTypeID    uint               `json:"exerciseTypeID,string"`
	SetGroupTemplates []SetGroupTemplate `json:"setgroupTemplates"`
	Metadata
}

type WorkoutTemplate struct {
	ID                uint               `json:"id,string"`
	UserID            uint               `json:"userID,string"`
	Name              string             `json:"name"`
	ExerciseTemplates []ExerciseTemplate `json:"exerciseTemplates"`
	Metadata
}
