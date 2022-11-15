package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/controller"
	"github.com/joeluhrman/Lift-Tracker/internal/pkg/model"
)

const (
	PORT    = ":8080"
	DB_PATH = "database.db"
)

func main() {
	err := model.InitDB(DB_PATH)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	controller.SetupEndpoints(engine)
	//engine.Use(static.Serve("/", static.LocalFile("./web/reactjs/build", true)))
	engine.Run(PORT)
}
