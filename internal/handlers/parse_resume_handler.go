package handlers

import (
	"log"
	"net/http"

	"github.com/Anilchoudary-Rugaramji/EZ-Apply/services"
	"github.com/gin-gonic/gin"
)

func ParseResumeHandler(c *gin.Context) {
	var requestBody struct {
		FileKey string `json:"file_key"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Invalid JSON request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	log.Println("Received request to parse resume. FileKey:", requestBody.FileKey)

	// Check if fileKey is empty
	if requestBody.FileKey == "" {
		log.Println("Error: file key is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "file_key cannot be empty"})
		return
	}

	parsedData, err := services.ProcessResume(requestBody.FileKey)
	if err != nil {
		log.Println("Error processing resume:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parsedData)
}
