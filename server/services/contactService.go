package services

import (
	"crypto/tls"
	"log"
	"net"
	"net/smtp"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type ContactService struct {
	secrets   *config.Secrets
	loginAuth utils.LoginAuth
}

func NewContactService(secrets *config.Secrets, loginAuth utils.LoginAuth) *ContactService {
	return &ContactService{
		secrets:   secrets,
		loginAuth: loginAuth,
	}
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
		return err
	}

	if err = c.Auth(service.loginAuth); err != nil {
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
