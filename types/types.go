package types

import (
	"errors"
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

func pplTypeFromString(s string) (pplType, error) {
	var pplTypes = []pplType{PPLTypePush, PPLTypePull, PPLTypeLegs}

	for _, pplType := range pplTypes {
		if s == string(pplType) {
			return pplType, nil
		}
	}

	return "", errors.New("Invalid PPLType: " + s)
}

func PPLTypesFromStrings(s []string) ([]pplType, error) {
	pplTypes := []pplType{}
	for _, str := range s {
		pplType, err := pplTypeFromString(str)
		if err != nil {
			return nil, err
		}

		pplTypes = append(pplTypes, pplType)
	}

	return pplTypes, nil
}

type musclegroup string

const (
	MscGrpCalves     musclegroup = "calves"
	MscGrpHamstrings musclegroup = "hamstrings"
	MscGrpQuads      musclegroup = "quads"
	MscGrpAdductors  musclegroup = "adductors"
	MscGrpAbductors  musclegroup = "abductors"
	MscGrpCore       musclegroup = "core"
	MscGrpLowBack    musclegroup = "lower back"
	MscGrpChest      musclegroup = "chest"
	MscGrpLats       musclegroup = "lats"
)

func musclegroupFromString(s string) musclegroup {
	var musclegroups = []musclegroup{
		MscGrpCalves, MscGrpHamstrings, MscGrpQuads,
		MscGrpAdductors, MscGrpAbductors, MscGrpCore,
		MscGrpLowBack, MscGrpChest, MscGrpLats,
	}

	for _, mscgrp := range musclegroups {
		if s == string(mscgrp) {
			return mscgrp
		}
	}

	return ""
}

func MuscleGroupsFromStrings(s []string) []musclegroup {
	mscgrps := []musclegroup{}
	for _, str := range s {
		mscgrp := musclegroupFromString(str)
		if mscgrp != "" {
			mscgrps = append(mscgrps, mscgrp)
		}
	}

	return mscgrps
}

type ExerciseType struct {
	ID           uint          `json:"id,string"`
	Name         string        `json:"name"`
	PPLTypes     []pplType     `json:"pplTypes"`
	MuscleGroups []musclegroup `json:"muscleGroups"`
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
	SetGroupLogs   []SetGroupLog `json:"setgroups"`
	metadata
}

type WorkoutLog struct {
	ID           uint          `json:"id,string"`
	UserID       uint          `json:"userID,string"`
	Date         time.Time     `json:"date"`
	Name         string        `json:"name"`
	Notes        string        `json:"notes"`
	ExerciseLogs []ExerciseLog `json:"exercises"`
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
	SetGroupTemplates []SetGroupTemplate `json:"setgroups"`
	metadata
}

type WorkoutTemplate struct {
	ID                uint               `json:"id,string"`
	UserID            uint               `json:"userID,string"`
	Name              string             `json:"name"`
	ExerciseTemplates []ExerciseTemplate `json:"exercises"`
	metadata
}
