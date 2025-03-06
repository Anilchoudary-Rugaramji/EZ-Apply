package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Anilchoudary-Rugaramji/EZ-Apply/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// create a gin router
	router := gin.Default()

	router.POST("/upload", handlers.UploadResume)

	// start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	router.Run(":" + port)
}
