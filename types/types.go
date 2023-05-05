package types

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type metadata struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
	IsAdmin        bool   `json:"isAdmin"`
	metadata
}

const SessionKey = "session"

type Session interface {
	Cookie() *http.Cookie
	Token() string
	UserID() uint
}

type session struct {
	userID uint
	token  string
	metadata
}

func NewSession(userID uint) *session {
	token := uuid.New().String()

	return &session{
		token:  token,
		userID: userID,
	}
}

func (s *session) Cookie() *http.Cookie {
	return &http.Cookie{
		Name:  SessionKey,
		Value: s.token,
	}
}

func (s *session) Token() string {
	return s.token
}

func (s *session) UserID() uint {
	return s.userID
}

type pplType string

const (
	PPLTypePush pplType = "push"
	PPLTypePull pplType = "pull"
	PPLTypeLegs pplType = "legs"
)

func pplTypeFromString(s string) pplType {
	var pplTypes = []pplType{PPLTypePush, PPLTypePull, PPLTypeLegs}

	for _, pplType := range pplTypes {
		if s == string(pplType) {
			return pplType
		}
	}

	return ""
}

func PPLTypesFromStrings(s []string) []pplType {
	pplTypes := []pplType{}
	for _, str := range s {
		pplTypes = append(pplTypes, pplTypeFromString(str))
	}

	return pplTypes
}

type MuscleGroup string

const (
	MscGrpCalves     MuscleGroup = "calves"
	MscGrpHamstrings MuscleGroup = "hamstrings"
	MscGrpQuads      MuscleGroup = "quads"
	MscGrpAdductors  MuscleGroup = "adductors"
	MscGrpAbductors  MuscleGroup = "abductors"
	MscGrpCore       MuscleGroup = "core"
	MscGrpLowBack    MuscleGroup = "lower back"
	MscGrpChest      MuscleGroup = "chest"
	MscGrpLats       MuscleGroup = "lats"
)

type ExerciseType struct {
	ID           uint          `json:"id,string"`
	Name         string        `json:"name"`
	PPLTypes     []pplType     `json:"pplTypes"`
	MuscleGroups []MuscleGroup `json:"muscleGroups"`
	metadata
}

type SetGroupLog struct {
	ID            uint    `json:"id,string"`
	ExerciseLogID uint    `json:"exerciseLogID,string"`
	Sets          uint    `json:"sets,string"`
	Reps          uint    `json:"reps,string"`
	Weight        float32 `json:"weight"`
	metadata
}

type ExerciseLog struct {
	ID             uint          `json:"id,string"`
	WorkoutLogID   uint          `json:"workoutLogID,string"`
	ExerciseTypeID uint          `json:"exerciseTypeID,string"`
	Notes          string        `json:"notes"`
	SetGroupLogs   []SetGroupLog `json:"setgroupLogs"`
	metadata
}

type WorkoutLog struct {
	ID           uint          `json:"id,string"`
	UserID       uint          `json:"userID,string"`
	Date         time.Time     `json:"date"`
	Name         string        `json:"name"`
	Notes        string        `json:"notes"`
	ExerciseLogs []ExerciseLog `json:"exerciseLogs"`
	metadata
}

type SetGroupTemplate struct {
	ID                 uint `json:"id,string"`
	ExerciseTemplateID uint `json:"exerciseTemplateID,string"`
	Sets               uint `json:"sets,string"`
	Reps               uint `json:"reps,string"`
	metadata
}

type ExerciseTemplate struct {
	ID                uint               `json:"id,string"`
	WorkoutTemplateID uint               `json:"workoutTemplateID,string"`
	ExerciseTypeID    uint               `json:"exerciseTypeID,string"`
	SetGroupTemplates []SetGroupTemplate `json:"setgroupTemplates"`
	metadata
}

type WorkoutTemplate struct {
	ID                uint               `json:"id,string"`
	UserID            uint               `json:"userID,string"`
	Name              string             `json:"name"`
	ExerciseTemplates []ExerciseTemplate `json:"exerciseTemplates"`
	metadata
}
