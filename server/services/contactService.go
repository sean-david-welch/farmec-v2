package services

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"io"
	"log"
	"net/smtp"

	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactService interface {
	SendContactEmail(data *types.EmailData) error
	ContactFormNotification(client *smtp.Client, data *types.EmailData) error
}

type ContactServiceImpl struct {
	secrets    *lib.Secrets
	smtpClient lib.SMTPClient
}

func NewContactService(secrets *lib.Secrets, smtpClient lib.SMTPClient) *ContactServiceImpl {
	return &ContactServiceImpl{
		secrets:    secrets,
		smtpClient: smtpClient,
	}
}

func (service *ContactServiceImpl) SendContactEmail(data *types.EmailData) error {
	client, err := service.smtpClient.SetupSMTPClient()
	if err != nil {
		log.Println("SMTP setup error:", err)
		return err
	}
	defer func(client *smtp.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	if err := service.ContactFormNotification(client, data); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}

func (service *ContactServiceImpl) ContactFormNotification(client *smtp.Client, data *types.EmailData) error {
	msg := fmt.Sprintf("Subject: New Contact Form From %s--%s\r\n\r\n%s", data.Name, data.Email, data.Message)

	if err := client.Mail(service.secrets.EmailUser); err != nil {
		return err
	}
	if err := client.Rcpt(service.secrets.EmailUser); err != nil {
		return err
	}

	wc, err := client.Data()
	if err != nil {
		return err
	}
	defer func(wc io.WriteCloser) {
		err := wc.Close()
		if err != nil {
			return
		}
	}(wc)

	if _, err := wc.Write([]byte(msg)); err != nil {
		return err
	}

	return nil
}
