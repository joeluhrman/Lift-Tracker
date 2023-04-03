// Contains types used across multiple packages.
package types

import (
	"image"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	IsAdmin        bool   `json:"is_admin"`

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

// incomplete
const (
	Calves     MuscleGroup = "calves"
	Hamstrings MuscleGroup = "hamstrings"
	Quads      MuscleGroup = "quads"
	Adductors  MuscleGroup = "adductors"
	Abductors  MuscleGroup = "abductors"
	Core       MuscleGroup = "core"
	LowBack    MuscleGroup = "lower back"
)

type ExerciseType struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Image       image.Image `json:"image";"omitempty"` // png only for now
	PPLType     PPLType     `json:"ppl_type"`
	MuscleGroup MuscleGroup `json:"muscle_group"`

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
	ExerciseLogID uint    `json:"exercise_log_id"`
	Sets          uint    `json:"sets"`
	Reps          uint    `json:"reps"`
	Weight        float32 `json:"weight"`

	Metadata
}

type ExerciseLog struct {
	ID             uint          `json:"id"`
	WorkoutLogID   uint          `json:"workout_log_id"`
	ExerciseTypeID uint          `json:"exercise_type_id"`
	Notes          string        `json:"notes"`
	SetGroupLogs   []SetGroupLog `json:"setgroup_logs"`

	Metadata
}

type WorkoutLog struct {
	ID           uint          `json:"id"`
	UserID       uint          `json:"user_id"`
	Date         time.Time     `json:"date"`
	Name         string        `json:"name"`
	Notes        string        `json:"notes"`
	ExerciseLogs []ExerciseLog `json:"exercise_logs"`

	Metadata
}

type SetGroupTemplate struct {
	ID                 uint `json:"id"`
	ExerciseTemplateID uint `json:"exercise_template_id"`
	Sets               uint `json:"sets"`
	Reps               uint `json:"reps"`

	Metadata
}

type ExerciseTemplate struct {
	ID                uint               `json:"id"`
	WorkoutTemplateID uint               `json:"workout_template_id"`
	ExerciseTypeID    uint               `json:"exercise_type_id"`
	SetGroupTemplates []SetGroupTemplate `json:"setgroup_templates"`

	Metadata
}

type WorkoutTemplate struct {
	ID                uint               `json:"id"`
	UserID            uint               `json:"user_id"`
	Name              string             `json:"name"`
	ExerciseTemplates []ExerciseTemplate `json:"exercise_templates"`

	Metadata
}
