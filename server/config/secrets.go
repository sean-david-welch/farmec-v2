package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Secrets struct {
	Domain string
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
	StripeSecretKey string
	StripePublicKey string

	EmailHost string
	EmailPass string
	EmailUser string
	EmailPort string
}

func NewSecrets() (*Secrets, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Secrets{
		Domain: os.Getenv("DOMAIN"),
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

		StripeSecretKey: os.Getenv("STRIPE_SECRET_KEY"),
		StripePublicKey: os.Getenv("STRIPE_PUBLIC_KEY"),

		EmailHost: os.Getenv("EMAIL_HOST"),
		EmailPass: os.Getenv("EMAIL_PASSWORD"),
		EmailUser: os.Getenv("EMAIL_USER"),
		EmailPort: os.Getenv("EMAIL_PASSWORD"),
	}, nil
}