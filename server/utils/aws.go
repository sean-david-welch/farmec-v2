package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/transport/http"
)

type S3Client interface {
	GeneratePresignedUrl(folder, image string) (string, string, error)
	DeleteImageFromS3(imageUrl string) error
}

type S3ClientImpl struct {
	Client *s3.Client
}

func NewS3Client(region, accessKey, secretKey string) (*S3ClientImpl, error) {
	conf, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)

	if err != nil {
		return nil, err
	}

	return &S3ClientImpl{
		Client: s3.NewFromConfig(conf),
	}, nil
}

func (client *S3ClientImpl) GeneratePresignedUrl(folder string, image string) (string, string, error) {
	const bucketName = "farmec.ie"
	const cloudfrontDomain = "https://farmec.ie"

	if folder == "" || image == "" {
		return "", "", fmt.Errorf("folder and image parameters must not be empty")
	}

	imageKey := fmt.Sprintf("farmec_images/%s/%s", folder, image)
	imageUrl := fmt.Sprintf("%s/%s", cloudfrontDomain, imageKey)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	presignClient := s3.NewPresignClient(client.Client)

	presignReq, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imageKey),
	}, s3.WithPresignExpires(5*time.Minute))

	if err != nil {
		var respErr *http.ResponseError
		if errors.As(err, &respErr) {
			return "", "", fmt.Errorf("AWS presign error: %s, StatusCode: %d", respErr.Error(), respErr.Response.StatusCode)
		}
		return "", "", fmt.Errorf("error generating presigned URL: %w", err)
	}

	return presignReq.URL, imageUrl, nil
}

func (client *S3ClientImpl) DeleteImageFromS3(imageUrl string) error {
	var respErr *http.ResponseError
	const bucketName = "farmec.ie"

	parsedUrl, err := url.Parse(imageUrl)
	if err != nil {
		return fmt.Errorf("error parsing URL: %w", err)
	}
	key := strings.TrimPrefix(parsedUrl.Path, "/")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = client.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		if errors.As(err, &respErr) {
			log.Printf("AWS error: %s, StatusCode: %d", respErr.Error(), respErr.HTTPStatusCode())
			return fmt.Errorf("AWS error: %s, StatusCode: %d", respErr.Error(), respErr.HTTPStatusCode())
		}
		log.Printf("Error deleting object from S3: %s, Object Key: %s", err.Error(), key)
		return fmt.Errorf("error deleting object from S3: %w", err)
	}

	return nil
}
