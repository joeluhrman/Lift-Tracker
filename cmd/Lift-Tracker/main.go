package main

import (
	"github.com/gin-gonic/gin"
)

const (
	PORT = ":8080"
)

func main() {
	engine := gin.Default()

	engine.Run()
}
