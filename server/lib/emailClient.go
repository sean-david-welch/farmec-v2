package lib

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

type EmailClient interface {
	SendFormNotification(data *types.EmailData, form string) error
}

type EmailClientImpl struct {
	secrets *Secrets
	client  *sendgrid.Client
}

func NewEmailClient(secrets *Secrets) *EmailClientImpl {
	client := sendgrid.NewSendClient(secrets.SendGridAPIKey)
	return &EmailClientImpl{
		secrets: secrets,
		client:  client,
	}
}

func (service *EmailClientImpl) SendFormNotification(data *types.EmailData, form string) error {
	from := mail.NewEmail("Farmec Ireland Ltd", service.secrets.EmailUser)
	to := mail.NewEmail("Admin", service.secrets.EmailUser)
	subject := fmt.Sprintf("New %s Form from %s--%s", form, data.Name, data.Email)

	plainContent := data.Message
	htmlContent := fmt.Sprintf("<p>%s</p>", data.Message)

	message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)

	log.Printf("Sending email with subject: %s", subject)
	response, err := service.client.Send(message)

	if err != nil {
		log.Printf("SendGrid ERROR: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully. StatusCode: %d", response.StatusCode)
	if response.Body != "" {
		log.Printf("Response body: %s", response.Body)
	}
	return nil
}
