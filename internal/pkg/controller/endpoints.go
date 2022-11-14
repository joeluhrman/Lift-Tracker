package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	END_API = "/api"
	END_V1  = "/v1"

	END_LOGIN = "/login"
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

		v1.GET(END_LOGIN, func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "test"})
		})
	}
}
