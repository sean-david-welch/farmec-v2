package lib

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailClient interface {
	SendFormNotification(data *types.EmailData, form string) error
}

type SendGridClient struct {
	secrets *Secrets
	client  *sendgrid.Client
}

func NewEmailClient(secrets *Secrets) *SendGridClient {
	client := sendgrid.NewSendClient(secrets.SendGridAPIKey)
	return &SendGridClient{
		secrets: secrets,
		client:  client,
	}
}

func (service *SendGridClient) SendFormNotification(data *types.EmailData, form string) error {
	from := mail.NewEmail("Farmec Website", service.secrets.EmailUser)
	to := mail.NewEmail("Admin", service.secrets.EmailUser)
	subject := fmt.Sprintf("New %s Form from %s--%s", form, data.Name, data.Email)

	plainContent := data.Message
	htmlContent := fmt.Sprintf("<p>%s</p>", data.Message)

	message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)

	_, err := service.client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
