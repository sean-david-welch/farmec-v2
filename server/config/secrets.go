package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Secrets struct {
	Domain                  string
	DatabaseURL             string
	DockerDatabase          string
	ProjectId               string
	PrivateKeyId            string
	PrivateKey              string
	ClientEmail             string
	ClientId                string
	AuthUri                 string
	TokenUri                string
	AuthProviderX509CertUrl string
	ClientX509CertUrl       string
	AwsAccessKey            string
	AwsSecret               string
	YoutubeApiKey           string
	StripeSecretKey         string
	StripePublicKey         string
	EmailHost               string
	EmailPass               string
	EmailUser               string
	EmailPort               string
}

func NewSecrets() (*Secrets, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Secrets{
		Domain: os.Getenv("DOMAIN"),
		// Database
		DatabaseURL: os.Getenv("DATABASE_URL"),
		// Docker DB
		DockerDatabase: os.Getenv("DOCKER_DATABASE"),
		// Firebase
		ProjectId:               os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyId:            os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("FIREBASE_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientId:                os.Getenv("FIREBASE_CLIENT_ID"),
		AuthUri:                 os.Getenv("FIREBASE_AUTH_URI"),
		TokenUri:                os.Getenv("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertUrl: os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertUrl:       os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
		// AWS
		AwsAccessKey: os.Getenv("AWS_ACCESS_KEY"),
		AwsSecret:    os.Getenv("AWS_SECRET_KEY"),
		// Youtube
		YoutubeApiKey: os.Getenv("YOUTUBE_API_KEY"),
		// Stripe
		StripeSecretKey: os.Getenv("STRIPE_SECRET_KEY"),
		StripePublicKey: os.Getenv("STRIPE_PUBLIC_KEY"),
		// Email
		EmailHost: os.Getenv("EMAIL_HOST"),
		EmailPass: os.Getenv("EMAIL_PASSWORD"),
		EmailUser: os.Getenv("EMAIL_USER"),
		EmailPort: os.Getenv("EMAIL_PASSWORD"),
	}, nil
}
