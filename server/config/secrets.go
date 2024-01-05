package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Secrets struct {
	DatabaseURL string
	ProjectId string
	PrivateKeyId string
	PrivateKey string
	ClientEmail string
	ClientId string
	AuthUri string
	TokenUri string
	AuthProviderX509CertUrl string
	ClientX509CertUrl string
	AwsAccessKey string
	AwsSecret string
	YoutubeApiKey string
}

func NewSecrets() (*Secrets, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Secrets{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ProjectId: os.Getenv("PROJECT_ID"),
		PrivateKeyId: os.Getenv("PRIVATE_KEY_ID"),
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		ClientEmail: os.Getenv("CLIENT_EMAIL"),
		ClientId: os.Getenv("CLIENT_ID"),
		AuthUri: os.Getenv("AUTH_URI"),
		TokenUri: os.Getenv("TOKEN_URI"),
		AuthProviderX509CertUrl: os.Getenv("AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertUrl: os.Getenv("CLIENT_X509_CERT_URL"),
		AwsAccessKey: os.Getenv("AWS_ACCESS_KEY"),
		AwsSecret: os.Getenv("AWS_SECRET"),
		YoutubeApiKey: os.Getenv("YOUTUBE_API_KEY"),
	}, nil
}