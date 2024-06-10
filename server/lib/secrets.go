package lib

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Secrets struct {
	DatabaseURL             string
	RailwayURL              string
	DockerDatabase          string
	RdsDatabase             string
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
	StripeSecretKeyTest     string
	EmailHost               string
	EmailPass               string
	EmailUser               string
	EmailPort               string
}

func NewSecrets() (*Secrets, error) {

	env := os.Getenv("ENV")
	if env != "production" {
		if err := godotenv.Load(".env"); err != nil {
			return nil, fmt.Errorf("failed to load configurations file: %w", err)
		}
	}

	return &Secrets{
		// Database
		DatabaseURL:    os.Getenv("RDS_DATABASE"),
		RailwayURL:     os.Getenv("RAILWAY_URL"),
		DockerDatabase: os.Getenv("DOCKER_DATABASE"),
		RdsDatabase:    os.Getenv("RDS_DATABASE"),
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
		StripeSecretKey:     os.Getenv("STRIPE_SECRET_KEY"),
		StripeSecretKeyTest: os.Getenv("TEST_SECRET_KEY"),
		// Email
		EmailHost: os.Getenv("EMAIL_HOST"),
		EmailPass: os.Getenv("EMAIL_PASSWORD"),
		EmailUser: os.Getenv("EMAIL_USER"),
		EmailPort: os.Getenv("EMAIL_PORT"),
	}, nil
}
