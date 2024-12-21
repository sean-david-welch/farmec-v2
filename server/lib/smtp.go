package lib

import (
	"crypto/tls"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"io"
	"net"
	"net/smtp"
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

func (service *SMTPClientImpl) SetupSMTPClient() (*smtp.Client, error) {
	host := "smtp.office365.com"
	port := 587
	addr := fmt.Sprintf("%s:%d", host, port)

	// Create auth mechanism
	auth := smtp.PlainAuth(
		"",                        // Identity (can be empty)
		service.secrets.EmailUser, // Username (full email)
		service.secrets.EmailPass, // App Password
		host,                      // Host for auth
	)

	// Connect first
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dial failed: %v", err)
	}

	client, err := smtp.NewClient(c, host)
	if err != nil {
		return nil, fmt.Errorf("client creation failed: %v", err)
	}

	// EHLO/HELO - use your domain
	if err = client.Hello("farmec.ie"); err != nil {
		return nil, fmt.Errorf("HELO failed: %v", err)
	}

	// Get server capabilities
	if ok, _ := client.Extension("STARTTLS"); ok {
		if err = client.StartTLS(&tls.Config{
			ServerName: host,
			MinVersion: tls.VersionTLS12,
		}); err != nil {
			return nil, fmt.Errorf("StartTLS failed: %v", err)
		}
	}

	// Get authentication mechanisms
	if ok, _ := client.Extension("AUTH"); ok {
		if err = client.Auth(auth); err != nil {
			return nil, fmt.Errorf("auth failed: %v", err)
		}
	}

	return client, nil
}

func (service *SMTPClientImpl) SendFormNotification(client *smtp.Client, data *types.EmailData, form string) error {
	msg := fmt.Sprintf("Subject: New %s Form from %s--%s\r\n\r\n%s", form, data.Name, data.Email, data.Message)

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
