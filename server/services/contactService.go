package services

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
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
	if err := service.smtpClient.SendFormNotification(data, "Contact"); err != nil {
		return err
	}
	return nil
}
