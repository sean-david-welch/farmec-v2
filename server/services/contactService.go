package services

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"net/smtp"

	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactService interface {
	SendContactEmail(data *types.EmailData) error
}

type ContactServiceImpl struct {
	smtpClient lib.SMTPClient
}

func NewContactService(smtpClient lib.SMTPClient) *ContactServiceImpl {
	return &ContactServiceImpl{
		smtpClient: smtpClient,
	}
}

func (service *ContactServiceImpl) SendContactEmail(data *types.EmailData) error {
	client, err := service.smtpClient.SetupSMTPClient()
	if err != nil {
		return err
	}
	defer func(client *smtp.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	if err := service.smtpClient.SendFormNotification(data, "Contact"); err != nil {
		return err
	}
	return nil
}
