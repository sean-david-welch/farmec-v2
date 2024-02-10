package services

import (
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/smtp"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactService struct {
	secrets *config.Secrets
}

func NewContactService(secrets *config.Secrets) *ContactService {
	return &ContactService{secrets: secrets}
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

func (service *ContactService) SendEmail(data *types.EmailData) error {
	conn, err := net.Dial("tcp", "smtp.office365.com:587")
	if err != nil {
		log.Println("Dial error:", err)
		return err
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, "smtp.office365.com")
	if err != nil {
		log.Println("SMTP client error:", err)
		return err
	}
	defer c.Close()

	tlsConfig := &tls.Config{ServerName: "smtp.office365.com"}
	if err = c.StartTLS(tlsConfig); err != nil {
		log.Println("StartTLS error:", err)
		return err
	}

	auth := LoginAuth(service.secrets.EmailUser, service.secrets.EmailPass)
	if err = c.Auth(auth); err != nil {
		log.Println("SMTP auth error:", err)
		return err
	}

	msg := "Subject: New Contact Form From " + data.Name + "--" + data.Email + "\r\n" +
		"\r\n" + data.Message

	if err = c.Mail(service.secrets.EmailUser); err != nil {
		log.Println("SMTP Mail error:", err)
		return err
	}
	if err = c.Rcpt(service.secrets.EmailUser); err != nil {
		log.Println("SMTP Rcpt error:", err)
		return err
	}

	wc, err := c.Data()
	if err != nil {
		log.Println("SMTP data error:", err)
		return err
	}
	defer wc.Close()
	if _, err = wc.Write([]byte(msg)); err != nil {
		log.Println("SMTP write error:", err)
		return err
	}

	return nil
}
