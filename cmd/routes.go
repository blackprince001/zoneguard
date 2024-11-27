package main

import (
	controllers "zoneguard/internal/controllers"
	"zoneguard/internal/database"

	"github.com/gin-gonic/gin"
)

type RoutesHandler struct {
	databaseRepo database.IplocatorRepo
}

func NewRoutesHandler(dbRepo database.IplocatorRepo) RoutesHandler {
	return RoutesHandler{
		databaseRepo: dbRepo,
	}
}

func (h *RoutesHandler) RegisterRoutes(r *gin.RouterGroup) {
	r = r.Group("")
	{
		r.GET("/", controllers.Root)
		r.GET("/address", controllers.Address)
		r.GET("/ip", controllers.GetIP)
		r.GET("/geo", controllers.GetIpGeoLocation)
	}
}
