package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"zoneguard/internal/database"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&database.Iplocator{})

	geoLocatorRepo := database.NewIplocatorRepo(db)
	routes_handler := NewRoutesHandler(geoLocatorRepo)

	csv_file := filepath.Join(".", "data.csv")

	err = database.CSVtoSqlite(db, csv_file)
	if err != nil {
		log.Fatal("Could not load all content of CSV to database")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	api := r.Group("/api")
	{
		routes_handler.RegisterRoutes(api)
	}
	r.Run(":8000")
}
