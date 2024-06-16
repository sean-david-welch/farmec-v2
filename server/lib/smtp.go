package lib

import (
	"crypto/tls"
	"net"
	"net/smtp"
)

type SMTPClient interface {
	SetupSMTPClient() (*smtp.Client, error)
}

type SMTPClientImpl struct {
	emailAuth EmailAuth
}

func (service *SMTPClientImpl) SetupSMTPClient() (*smtp.Client, error) {
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

	if err = client.Auth(service.emailAuth); err != nil {
		err := client.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return client, nil
}
