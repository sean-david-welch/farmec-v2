package secrets

import (
	"os"

	"github.com/joho/godotenv"
)

type Secrets struct {
	DatabaseURL string
}

func NewSecrets() (*Secrets, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Secrets{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}, nil
}