package storage

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadFile uploads a file to S3 and returns the URL to the file
func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	bucket := os.Getenv("AWS_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		log.Println("Failed to create session", err)
		return "", err
	}

	// Read file content into a buffer
	buffer := bytes.NewBuffer(nil)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		log.Println("Failed to read file", err)
		return "", err
	}

	// Genarate unique filename
	fileName := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), fileHeader.Filename)

	s3Client := s3.New(sess)
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: aws.String(http.DetectContentType(buffer.Bytes())),
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		log.Println("Failed to upload file", err)
		return "", err
	}

	fileUrl := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucket, region, fileName)
	return fileUrl, nil

}
