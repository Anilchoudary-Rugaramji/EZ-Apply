package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Anilchoudary-Rugaramji/EZ-Apply/internal/storage"
)

func ProcessResume(fiileKey string) (map[string]interface{}, error) {

	log.Println("Fetching file from S3 with key:", fiileKey)
	if fiileKey == "" {
		return nil, errors.New("file key is empty")
	}

	fileURL, err := storage.GetFileFromS3(fiileKey)
	if err != nil {
		log.Println("Error fetching file from S3:", err)
		return nil, err
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		log.Println("Error downloading file from S3:", err)
		return nil, err
	}
	defer resp.Body.Close()

	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading file content:", err)
		return nil, err
	}

	fastAPIURL := "http://127.0.0.1:8000/parse_pdf"
	req, err := http.NewRequest("POST", fastAPIURL, bytes.NewBuffer(fileBytes))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/pdf")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Println("Error sending request to FastAPI:", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response data:", err)
		return nil, err
	}

	var parsedData map[string]interface{}
	if err := json.Unmarshal(responseData, &parsedData); err != nil {
		log.Println("Error parsing JSON response:", err)
		return nil, err
	}
	fmt.Println("final data before going to postgres:", parsedData)
	return parsedData, nil
}
