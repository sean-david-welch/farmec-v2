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
	conn, err := net.Dial("tcp", "smtp-legacy.office365.com:587")
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, "smtp-legacy.office365.com")
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	tlsConfig := &tls.Config{ServerName: "smtp-legacy.office365.com"}
	if err = client.StartTLS(tlsConfig); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	if err = client.Auth(service.emailAuth); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
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
