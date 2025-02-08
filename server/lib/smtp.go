package lib

import (
	"crypto/tls"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"io"
	"net"
	"net/smtp"
	"time"
)

type SMTPClient interface {
	SetupSMTPClient() (*smtp.Client, error)
	SendFormNotification(data *types.EmailData, form string) error
}

type SMTPClientImpl struct {
	secrets   *Secrets
	emailAuth smtp.Auth
}

func NewSTMPClient(secrets *Secrets) *SMTPClientImpl {
	auth := NewOffice365Auth(
		secrets.EmailUser,
		secrets.EmailPass,
	)

	return &SMTPClientImpl{
		secrets:   secrets,
		emailAuth: auth,
	}
}

func (service *SMTPClientImpl) SetupSMTPClient() (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", "smtp-legacy.office365.com:587", 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to establish TCP connection: %w", err)
	}

	client, err := smtp.NewClient(conn, "smtp-legacy.office365.com")
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to create SMTP client: %w", err)
	}

	if err = client.Hello("localhost"); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("HELO failed: %w", err)
	}

	tlsConfig := &tls.Config{
		ServerName: "smtp-legacy.office365.com",
		MinVersion: tls.VersionTLS12,
	}

	if err = client.StartTLS(tlsConfig); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("TLS startup failed: %w", err)
	}

	if err = client.Auth(service.emailAuth); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	return client, nil
}

func (service *SMTPClientImpl) SendFormNotification(data *types.EmailData, form string) error {
	client, err := service.SetupSMTPClient()
	if err != nil {
		return fmt.Errorf("failed to setup client: %w", err)
	}
	defer func(client *smtp.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	msg := fmt.Sprintf("Subject: New %s Form from %s--%s\r\n\r\n%s",
		form, data.Name, data.Email, data.Message)

	if err := client.Mail(service.secrets.EmailUser); err != nil {
		return fmt.Errorf("MAIL FROM failed: %w", err)
	}

	if err := client.Rcpt(service.secrets.EmailUser); err != nil {
		return fmt.Errorf("RCPT TO failed: %w", err)
	}

	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("DATA command failed: %w", err)
	}
	defer func(wc io.WriteCloser) {
		err := wc.Close()
		if err != nil {
			return
		}
	}(wc)

	if _, err = wc.Write([]byte(msg)); err != nil {
		return fmt.Errorf("write failed: %w", err)
	}

	return nil
}
