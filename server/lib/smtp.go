package lib

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net/mail"
	"net/smtp"
	"strings"
)

type SMTPClientImpl struct {
	secrets *Secrets
}

func NewSTMPClient(secrets *Secrets) *SMTPClientImpl {
	return &SMTPClientImpl{secrets: secrets}
}

func (service *SMTPClientImpl) SendEmail(to []string, subject, body string) error {
	smtpHost := "smtp.office365.com"
	smtpPort := 587

	// Message composition
	from := mail.Address{Name: "Farmec", Address: service.secrets.EmailUser}
	message := strings.Builder{}
	message.WriteString(fmt.Sprintf("From: %s\r\n", from.String()))
	message.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ", "))) // Changed separator to comma+space
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	message.WriteString("MIME-Version: 1.0\r\n")
	message.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	message.WriteString("\r\n")
	message.WriteString(body)

	// Use PlainAuth with App Password
	auth := smtp.PlainAuth(
		"",                        // Identity
		service.secrets.EmailUser, // Username (full email address)
		service.secrets.EmailPass, // App Password (NOT your regular password)
		smtpHost,
	)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpHost, smtpPort),
		auth,
		service.secrets.EmailUser,
		to,
		[]byte(message.String()),
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func (service *SMTPClientImpl) SendFormNotification(data *types.EmailData, form string) error {
	to := []string{service.secrets.EmailUser}
	subject := fmt.Sprintf("New %s Form from %s", form, data.Name)
	body := fmt.Sprintf("From: %s\nEmail: %s\n\nMessage:\n%s", data.Name, data.Email, data.Message)

	return service.SendEmail(to, subject, body)
}
