package lib

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net/mail"
	"net/smtp"
	"strings"
)

type SMTPClient interface {
	SetupSMTPClient() (*smtp.Client, error)
	SendFormNotification(client *smtp.Client, data *types.EmailData, form string) error
}

type SMTPClientImpl struct {
	secrets   *Secrets
	emailAuth EmailAuth
}

func NewSTMPClient(secrets *Secrets, emailAuth EmailAuth) *SMTPClientImpl {
	return &SMTPClientImpl{secrets: secrets, emailAuth: emailAuth}
}

func (service *SMTPClientImpl) SendEmail(to []string, subject, body string) error {
	smtpHost := "smtp.office365.com"
	smtpPort := 587

	fmt.Printf("Attempting to send email via %s:%d", smtpHost, smtpPort)
	fmt.Printf("From: %s", service.secrets.EmailUser)
	fmt.Printf("To: %s", strings.Join(to, ";"))
	fmt.Printf("Subject: %s", subject)

	// Message composition
	from := mail.Address{Name: "Farmec", Address: service.secrets.EmailUser}
	message := strings.Builder{}
	message.WriteString(fmt.Sprintf("From: %s\r\n", from.String()))
	message.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ";")))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	message.WriteString("\r\n")
	message.WriteString(body)

	// Use CRAM-MD5 auth instead of PlainAuth
	auth := smtp.CRAMMD5Auth(service.secrets.EmailUser, service.secrets.EmailPass)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpHost, smtpPort),
		auth,
		service.secrets.EmailUser,
		to,
		[]byte(message.String()),
	)

	if err != nil {
		fmt.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Printf("Email sent successfully")
	return nil
}

func (service *SMTPClientImpl) SendFormNotification(data *types.EmailData, form string) error {
	fmt.Printf("Sending form notification for %s form", form)

	to := []string{service.secrets.EmailUser}
	subject := fmt.Sprintf("New %s Form from %s", form, data.Name)
	body := fmt.Sprintf("From: %s\nEmail: %s\n\nMessage:\n%s", data.Name, data.Email, data.Message)

	return service.SendEmail(to, subject, body)
}
