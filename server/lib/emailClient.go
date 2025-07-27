package lib

import (
	"fmt"
	"github.com/resend/resend-go/v2"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
)

type EmailClient interface {
	SendFormNotification(data *types.EmailData, form string) error
}

type EmailClientImpl struct {
	secrets *Secrets
	client  *resend.Client
}

func NewEmailClient(secrets *Secrets) *EmailClientImpl {
	client := resend.NewClient(secrets.ResendToken)
	return &EmailClientImpl{
		secrets: secrets,
		client:  client,
	}
}

func (service *EmailClientImpl) SendFormNotification(data *types.EmailData, form string) error {
	subject := fmt.Sprintf("New %s Form from %s--%s", form, data.Name, data.Email)
	plainContent := data.Message
	htmlContent := fmt.Sprintf("<p>%s</p>", data.Message)

	params := &resend.SendEmailRequest{
		From:    "Farmec Ireland Ltd <noreply@farmec.ie>",
		To:      []string{service.secrets.EmailUser},
		Subject: subject,
		Text:    plainContent,
		Html:    htmlContent,
	}

	log.Printf("Sending email with subject: %s", subject)
	sent, err := service.client.Emails.Send(params)

	if err != nil {
		log.Printf("Resend ERROR: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully. ID: %s", sent.Id)
	return nil
}
