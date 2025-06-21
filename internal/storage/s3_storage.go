package storage

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadFile uploads a file to S3 and returns the URL
func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	bucket := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	// Initialize AWS Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_KEY"),
			"",
		),
	})
	if err != nil {
		log.Println("Failed to create AWS session:", err)
		return "", "", err
	}

	// Read file into a buffer
	buffer := bytes.NewBuffer(nil)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		log.Println("Failed to read file:", err)
		return "", "", err
	}

	// Generate a unique file key (This is needed later)
	fileKey := fmt.Sprintf("%s-%d-%s", uuid.New().String(), time.Now().Unix(), fileHeader.Filename)

	// Create S3 Client
	s3Client := s3.New(sess)

	// Upload to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileKey),
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})
	if err != nil {
		log.Println("Failed to upload file:", err)
		return "", "", err
	}

	// Construct the S3 file URL
	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileKey)
	return fileURL, fileKey, nil
}

func GetFileFromS3(fileKey string) (string, error) {
	bucket := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	// Initialize AWS Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_KEY"),
			"",
		),
	})
	if err != nil {
		log.Println("Failed to create AWS session:", err)
		return "", err
	}

	s3Client := s3.New(sess)

	// Generate a pre-signed URL (Valid for 1 Hour)
	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})

	presignedURL, err := req.Presign(1 * time.Hour)
	if err != nil {
		log.Println("Failed to generate pre-signed URL:", err)
		return "", err
	}

	fmt.Println("Generated Pre-Signed URL:", presignedURL) // Debugging
	return presignedURL, nil
}
