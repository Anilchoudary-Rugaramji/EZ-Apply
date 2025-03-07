package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anilchoudary-Rugaramji/EZ-Apply/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func UplaoadResume(c *gin.Context) {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	// Get the uploaded file
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Upload file to S3
	fileURL, err := storage.UploadFile(file, fileHeader)
	if err != nil {
		fmt.Println("S3 Upload Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Return file URL
	c.JSON(http.StatusOK, gin.H{
		"message":  "Resume uploaded successfully",
		"file_url": fileURL,
	})
}
