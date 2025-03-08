package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Anilchoudary-Rugaramji/EZ-Apply/internal/handlers"
	"github.com/Anilchoudary-Rugaramji/EZ-Apply/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Connect to the database
	db, err := storage.ConnectDataBase()
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	// Migrate the database (ensure tables are created)
	err = storage.MigrateDB(db)
	if err != nil {
		log.Fatal("Error migrating database", err)
	}

	// Create a Gin router
	router := gin.Default()

	// Optionally, you can enable CORS middleware if needed
	// router.Use(cors.Default())

	// Define route and associate it with handler
	router.POST("/upload", handlers.UploadResume)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	router.Run(":" + port)
}
