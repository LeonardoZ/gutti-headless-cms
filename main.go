package main

import (
	"log"
	"os"

	"github.com/LeonardoZ/gutti-headless-cms/config"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// setup DB
	_, err := config.ConnectDb()
	if err != nil {
		log.Fatal("Cannot connect to Database")
	}
	// setup router
	router := gin.Default()
	port := os.Getenv("API_PORT")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port)
}
