package services

import (
	"crypto/tls"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"io"
	"log"
	"net"
	"net/smtp"

	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactService interface {
	SendEmail(data *types.EmailData) error
	SetupSMTPClient() (*smtp.Client, error)
	SendMessage(client *smtp.Client, data *types.EmailData) error
}

type ContactServiceImpl struct {
	secrets   *lib.Secrets
	loginAuth lib.EmailAuth
}

func NewContactService(secrets *lib.Secrets, loginAuth lib.EmailAuth) *ContactServiceImpl {
	return &ContactServiceImpl{
		secrets:   secrets,
		loginAuth: loginAuth,
	}
}

func (service *ContactServiceImpl) SendEmail(data *types.EmailData) error {
	client, err := service.SetupSMTPClient()
	if err != nil {
		log.Println("SMTP setup error:", err)
		return err
	}
	defer func(client *smtp.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	if err := service.SendMessage(client, data); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}

func (service *ContactServiceImpl) SetupSMTPClient() (*smtp.Client, error) {
	conn, err := net.Dial("tcp", "smtp.office365.com:587")
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, "smtp.office365.com")
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	tlsConfig := &tls.Config{ServerName: "smtp.office365.com"}
	if err = client.StartTLS(tlsConfig); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	if err = client.Auth(service.loginAuth); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return client, nil
}

func (service *ContactServiceImpl) SendMessage(client *smtp.Client, data *types.EmailData) error {
	msg := fmt.Sprintf("Subject: New Contact Form From %s--%s\r\n\r\n%s", data.Name, data.Email, data.Message)

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
