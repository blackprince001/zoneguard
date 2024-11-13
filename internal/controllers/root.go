package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "ZoneGuard",
	})
}
