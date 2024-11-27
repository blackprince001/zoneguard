package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"zoneguard/internal/database"
	"zoneguard/internal/grubber"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetIP(c *gin.Context) {
	ip, err := grubber.GetRemoteIP(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    ip,
	})
}

func GetIpGeoLocation(c *gin.Context) {
	ip, err := grubber.GetRemoteIP(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var geoData database.Iplocator

	related_id := strings.Join(strings.Split(ip.IpAddr, "."), "")
	relatedID, err := strconv.ParseUint(related_id, 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = db.Where("related_id = ?", relatedID).Find(&geoData).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    ip,
		"geo":     geoData,
	})
}

func Address(c *gin.Context) {
	ip, err := grubber.GetRemoteIP(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(200, "ip.tmpl", gin.H{
		"ip_address": ip.IpAddr,
	})
}
