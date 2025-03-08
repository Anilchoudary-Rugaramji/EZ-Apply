package handlers

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func getMetadataAndUploadtoDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	// Get the uploaded file from the s3

	// send the api request to the python script to parse pdf to the json

	// save t he json data to the post gres table ...

}

func GetFileFromS3(fileKey string) (string, error) {
	bucket := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_KEY"),
			"",
		),
	})
	if err != nil {
		log.Println("Failed to create AWS session", err)
		return "", err
	}

	s3Client := s3.New(sess)

	// Generate a pre-signed URL valid for 1 hour
	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})

	presignedURL, err := req.Presign(1 * time.Hour)
	if err != nil {
		log.Println("Failed to generate pre-signed URL", err)
		return "", err
	}

	return presignedURL, nil
}
