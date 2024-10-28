package lib

import (
	"crypto/tls"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net"
	"net/smtp"
	"time"
)

type SMTPClient interface {
	SetupSMTPClient() (*smtp.Client, error)
	SendFormNotification(client *smtp.Client, data *types.EmailData, form string) error
}

type SMTPClientImpl struct {
	secrets   *Secrets
	emailAuth EmailAuth
}

func NewSTMPClient(secrets *Secrets) *SMTPClientImpl {
	auth := NewLoginAuth(
		secrets.EmailUser, // Should be full email address
		secrets.EmailPass, // Should be app password
	)
	return &SMTPClientImpl{
		secrets:   secrets,
		emailAuth: auth,
	}
}

func (service *SMTPClientImpl) SetupSMTPClient() (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", "smtp.office365.com:587", 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to establish TCP connection: %w", err)
	}

	client, err := smtp.NewClient(conn, "smtp.office365.com")
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to create SMTP client: %w", err)
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

	tlsConfig := &tls.Config{ServerName: "smtp.office365.com", MinVersion: tls.VersionTLS12}
	if err = client.StartTLS(tlsConfig); err != nil {
		err := client.Close()
		if err != nil {
			return nil, fmt.Errorf("TLS startup failed: %w", err)

		}
		return nil, fmt.Errorf("TLS startup failed: %w", err)
	}

	if err = client.Auth(service.emailAuth); err != nil {
		err := client.Close()
		if err != nil {
			return nil, fmt.Errorf("authentication failed: %w", err)
		}
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	return client, nil
}

func (service *SMTPClientImpl) SendFormNotification(data *types.EmailData, form string) error {
	// Create new client for each send operation
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

	_, err = wc.Write([]byte(msg))
	if err != nil {
		err := wc.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("write failed: %w", err)
	}

	err = wc.Close()
	if err != nil {
		return fmt.Errorf("close failed: %w", err)
	}

	return nil
}
