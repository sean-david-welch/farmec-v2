package services

import (
	"net/smtp"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactService struct {
	secrets *config.Secrets
}

func NewContactService(secrets *config.Secrets) *ContactService {
	return &ContactService{secrets: secrets}
}

func(service *ContactService) SendEmail(data *types.EmailData) error {
	auth := smtp.PlainAuth()
}