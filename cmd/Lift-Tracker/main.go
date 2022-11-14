package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/controller"
)

const (
	PORT = ":8080"
)

func main() {
	engine := gin.Default()
	controller.SetupEndpoints(engine)
	engine.Use(static.Serve("/", static.LocalFile("./web/reactjs/build", true)))
	engine.Run(PORT)
}
