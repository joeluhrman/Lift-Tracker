package controller

import (
	"bytes"
	"database/sql"
	"net/http"
	"testing"

	"github.com/joeluhrman/Lift-Tracker/internal/pkg/controller/handlers"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/model"
)

const (
	CODE_ERROR = "code was %d, should have been %d"
)

// Test correct response code when everything is valid.
func Test_EndPostCreateAccount(t *testing.T) {
	defer cleanUpDB()

	const USERNAME = "jaluhrman"
	const PASSWORD = "password"

	engine := setupTestEngine()

	postForm := []byte("username=" + USERNAME + "&password=" + PASSWORD)

	w := sendMockHTTPRequest(http.MethodPost, END_API+END_V1+END_POST_CREATE_ACCOUNT,
		bytes.NewBuffer(postForm), CON_TYPE_PF, engine)

	if w.Code != handlers.CREATE_ACC_SUCCESS_CODE {
		t.Errorf(CODE_ERROR, w.Code, handlers.CREATE_ACC_SUCCESS_CODE)
	}
}

// Test correct response code when username does not meet requirements.
func Test_EndPostCreateAccount_NoUsername(t *testing.T) {
	defer cleanUpDB()

	const USERNAME = ""
	const PASSWORD = "password"

	engine := setupTestEngine()

	postForm := []byte("username=" + USERNAME + "&password=" + PASSWORD)

	w := sendMockHTTPRequest(http.MethodPost, END_API+END_V1+END_POST_CREATE_ACCOUNT,
		bytes.NewBuffer(postForm), CON_TYPE_PF, engine)

	if w.Code != handlers.CREATE_ACC_CRED_REQS_NOT_MET {
		t.Errorf(CODE_ERROR, w.Code, handlers.CREATE_ACC_CRED_REQS_NOT_MET)
	}
}

// Test correct response code when password does not meet requirements.
func Test_EndPostCreateAccount_NoPassword(t *testing.T) {
	defer cleanUpDB()

	const USERNAME = "jaluhrman"
	const PASSWORD = ""

	engine := setupTestEngine()

	postForm := []byte("username=" + USERNAME + "&password=" + PASSWORD)

	w := sendMockHTTPRequest(http.MethodPost, END_API+END_V1+END_POST_CREATE_ACCOUNT,
		bytes.NewBuffer(postForm), CON_TYPE_PF, engine)

	if w.Code != handlers.CREATE_ACC_CRED_REQS_NOT_MET {
		t.Errorf(CODE_ERROR, w.Code, handlers.CREATE_ACC_CRED_REQS_NOT_MET)
	}
}

// Test correct response code when user already exists.
func Test_EndPostCreateAccount_AlreadyExists(t *testing.T) {
	defer cleanUpDB()

	model.DBConn.Create(&model.User{
		Username: sql.NullString{
			String: "jaluhrman",
			Valid:  true,
		},
	})

	const USERNAME = "jaluhrman"
	const PASSWORD = "password"

	engine := setupTestEngine()

	postForm := []byte("username=" + USERNAME + "&password=" + PASSWORD)

	w := sendMockHTTPRequest(http.MethodPost, END_API+END_V1+END_POST_CREATE_ACCOUNT,
		bytes.NewBuffer(postForm), CON_TYPE_PF, engine)

	if w.Code != handlers.CREATE_ACC_ALREADY_EXISTS {
		t.Errorf(CODE_ERROR, w.Code, handlers.CREATE_ACC_ALREADY_EXISTS)
	}
}
