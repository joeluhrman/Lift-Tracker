package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/controller"
)

const (
	PORT = ":8080"
)

func main() {
	engine := gin.Default()
	controller.SetupEndpoints(engine)
	engine.Run(PORT)
}
