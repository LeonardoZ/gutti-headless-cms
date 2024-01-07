package main

import (
	"log"
	"os"

	"github.com/LeonardoZ/gutti-headless-cms/config"
	"github.com/LeonardoZ/gutti-headless-cms/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// setup DB
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatal("Cannot connect to Database")
	}
	port := os.Getenv("API_PORT")

	// setup router
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	controller.RegisterContentBlockRoutes(db, router)
	router.Run(":" + port)
}
