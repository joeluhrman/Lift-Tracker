package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/controller/cred_reqs"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/model"
)

const (
	CREATE_ACC_SUCCESS_CODE      = http.StatusCreated
	CREATE_ACC_CRED_REQS_NOT_MET = http.StatusForbidden
	CREATE_ACC_ALREADY_EXISTS    = http.StatusConflict
)

// Receives username and password in Post Form and creates user through the model.
func CreateAccount(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	err := cred_reqs.CheckUsernameMeetsRequirements(username)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(CREATE_ACC_CRED_REQS_NOT_MET, gin.H{
			"error": err,
		})
		return
	}

	err = cred_reqs.CheckPasswordMeetsRequirements(password)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(CREATE_ACC_CRED_REQS_NOT_MET, gin.H{
			"error": err,
		})
		return
	}

	user := &model.User{
		Username: sql.NullString{
			String: username,
			Valid:  true,
		},
	}

	userPassword := &model.UserPassword{
		Password: sql.NullString{
			String: password,
			Valid:  true,
		},
	}

	err = model.CreateUser(user, userPassword)
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}

	ctx.Status(CREATE_ACC_SUCCESS_CODE)
}

// placeholder
func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "ayy yuh",
	})
}
