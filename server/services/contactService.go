package services

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactService interface {
	SendContactEmail(data *types.EmailData) error
}

type ContactServiceImpl struct {
	smtpClient *lib.SMTPClientImpl
}

func NewContactService(smtpClient *lib.SMTPClientImpl) *ContactServiceImpl {
	return &ContactServiceImpl{
		smtpClient: smtpClient,
	}
}

func (service *ContactServiceImpl) SendContactEmail(data *types.EmailData) error {
	go func() {
		err := service.smtpClient.SendFormNotification(data, "Contact")
		if err != nil {
			return
		}
	}()
	return nil
}
