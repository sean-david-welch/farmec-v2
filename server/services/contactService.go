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
	auth := smtp.PlainAuth("", service.secrets.EmailUser, service.secrets.EmailPass, service.secrets.EmailHost)

	to := []string{service.secrets.EmailUser}
	msg := []byte("To: " + service.secrets.EmailUser + "\r\n" +
	"Subject: New Contact Form sent from: " + data.Name + " - " + data.Email + "\r\n" +
	"\r\n" +
	data.Message + "\r\n")


	if err := smtp.SendMail(service.secrets.EmailHost+":"+service.secrets.EmailPort, auth, data.Email, to, msg); err != nil {
		return err
	}

    return nil
}