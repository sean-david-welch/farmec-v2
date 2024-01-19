package utils

import (
	"context"
	"fmt"
	"net/url"
	"path"
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
    const cloudfrontDomain = "https://d3eerclezczw8.cloudfront.net"
    
    imageKey := fmt.Sprintf("%s/%s", folder, image)
    imageUrl := fmt.Sprintf("%s/%s", cloudfrontDomain, imageKey)

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

func (client *S3Client) DeleteImageFromS3(imageUrl string) error {
    const bucketName = "farmec-bucket"

    parsedUrl, err := url.Parse(imageUrl); if err != nil {
        return fmt.Errorf("error parsing url: %w", err)
    }

    key := path.Base(parsedUrl.Path)

    _, err = client.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
    })

    return err
}
