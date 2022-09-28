package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamowei/senv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/check/profile", func(c *gin.Context) {
		c.String(200, "config.name")
	})
	return r
}

func main() {
	appName := "config-client"
	config := senv.NewConfig(
		"localhost",
		"8888",
		appName,
		[]string{"dev"},
		"master")
	config.Fetch(true, true)
	config.Process()

	r := setupRouter()
	r.Run(":8080")
}
