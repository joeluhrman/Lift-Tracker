package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/controller/handlers"
)

const (
	END_API = "/api"
	END_V1  = "/v1"

	END_POST_CREATE_ACCOUNT = "/create-account"
	END_POST_LOGIN          = "/login"
)

func SetupEndpoints(engine *gin.Engine) {
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": "PAGE_NOT_FOUND", "message": "Page not found",
		})
	})

	api := engine.Group(END_API)
	{
		v1 := api.Group(END_V1)

		v1.POST(END_POST_CREATE_ACCOUNT, handlers.CreateAccount)
		v1.POST(END_POST_LOGIN, handlers.Login)
	}
}
