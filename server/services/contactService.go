package services

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactService interface {
	SendContactEmail(data *types.EmailData) error
}

type ContactServiceImpl struct {
	emialClient *lib.EmailClientImpl
}

func NewContactService(emialClient *lib.EmailClientImpl) *ContactServiceImpl {
	return &ContactServiceImpl{
		emialClient: emialClient,
	}
}

func (service *ContactServiceImpl) SendContactEmail(data *types.EmailData) error {
	fmt.Println("ContactService: Starting email send")
	err := service.emialClient.SendFormNotification(data, "Contact")
	if err != nil {
		fmt.Printf("ContactService: Failed to send email: %v\n", err)
		return err
	}
	fmt.Println("ContactService: Email sent successfully")
	return nil
}
