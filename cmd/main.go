package main

import (
	"github.com/gin-gonic/gin"

	controllers "zoneguard/internal/controllers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", controllers.Root)
	r.GET("/address", controllers.Address)
	r.GET("/ip", controllers.GetIP)

	r.Run(":8000")
}
