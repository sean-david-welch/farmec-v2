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
	fmt.Println("=== START EMAIL SENDING ===")
	fmt.Printf("Form type: %s\n", form)
	fmt.Printf("From email: %s\n", service.secrets.EmailUser)
	fmt.Printf("To email: %s\n", service.secrets.EmailUser)
	fmt.Printf("Sender name: %s\n", data.Name)
	fmt.Printf("Sender email: %s\n", data.Email)

	from := mail.NewEmail("Farmec Ireland Ltd", service.secrets.EmailUser)
	to := mail.NewEmail("Admin", service.secrets.EmailUser)
	subject := fmt.Sprintf("New %s Form from %s--%s", form, data.Name, data.Email)

	plainContent := data.Message
	htmlContent := fmt.Sprintf("<p>%s</p>", data.Message)

	message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)

	fmt.Printf("Sending email with subject: %s\n", subject)
	response, err := service.client.Send(message)

	if err != nil {
		fmt.Printf("SendGrid ERROR: %v\n", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Printf("Email sent successfully. StatusCode: %d\n", response.StatusCode)
	if response.Body != "" {
		fmt.Printf("Response body: %s\n", response.Body)
	}
	fmt.Println("=== END EMAIL SENDING ===")

	return nil
}
