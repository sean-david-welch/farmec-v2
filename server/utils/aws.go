package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
    Client *s3.Client
}

func NewS3Client(region, accessKey, secretKey string) (*S3Client, error) {
    conf, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(region),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
    )

    if err != nil {
        return nil, err
    }

    return &S3Client{
        Client: s3.NewFromConfig(conf),
    }, nil
}

func (client *S3Client) GeneratePresignedUrl(folder string, image string) (string, string, error) {
    const bucketName = "farmec-bucket"
    
    imageKey := fmt.Sprintf("%s/%s", folder, image)
    imageUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, imageKey)

    presignClient := s3.NewPresignClient(client.Client)


    presignReq, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(imageKey),
    }, s3.WithPresignExpires(5*time.Minute))

    if err != nil {
        return "", "", err
    }

    return presignReq.URL, imageUrl, nil
}

func (client *S3Client) GenerateDeletePresignedUrl(bucketName, key string) (string, error) {
    presignClient := s3.NewPresignClient(client.Client)
    presignReq, err := presignClient.PresignDeleteObject(context.TODO(), &s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
    }, s3.WithPresignExpires(5*time.Minute))

    if err != nil {
        return "", err
    }

    return presignReq.URL, nil
}
