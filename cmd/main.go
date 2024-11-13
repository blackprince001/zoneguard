package main

import (
	"github.com/gin-gonic/gin"

	controllers "zoneguard/internal/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/ip", controllers.GetIP)

	r.Run(":8080")
}
