package controllers

import (
	"net/http"
	"zoneguard/internal/grubber"

	"github.com/gin-gonic/gin"
)

func GetIP(c *gin.Context) {
	ip, err := grubber.GetRemoteIP(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, ip)
}
