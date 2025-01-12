package services

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type FileService struct{}

func (myf *FileService) UploadUsingS3(fileContent *bytes.Reader, fileName string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		return "", err
	}

	// Define AWS credentials and bucket information
	awsAccessKeyID := os.Getenv("LIARA_ACCESS_KEY")
	awsSecretAccessKey := os.Getenv("LIARA_SECRET_KEY")
	endpoint := os.Getenv("LIARA_ENDPOINT")
	bucketName := os.Getenv("LIARA_BUCKET_NAME")

	// Initialize S3 client with custom configuration
	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     awsAccessKeyID,
			SecretAccessKey: awsSecretAccessKey,
		}, nil
	})

	cfg.BaseEndpoint = aws.String(endpoint)

	client := s3.NewFromConfig(cfg)
	// Specify the destination key in the bucket
	destinationKey := "uploads/" + fileName

	// Use the S3 client to upload the file
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(destinationKey),
		Body:   fileContent,
	})

	baseUrl := "https://carizma-motors.storage.c2.liara.space/" + destinationKey

	return baseUrl, err
}
