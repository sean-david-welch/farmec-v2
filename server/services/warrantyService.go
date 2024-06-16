package services

import (
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"io"
	"net/smtp"
)

type WarrantyService interface {
	GetWarranties() ([]types.DealerOwnerInfo, error)
	GetWarrantyById(id string) (*types.WarrantyClaim, []types.PartsRequired, error)
	CreateWarranty(warranty *types.WarrantyClaim, parts []types.PartsRequired) error
	UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error
	DeleteWarranty(id string) error
}

type WarrantyServiceImpl struct {
	secrets    *lib.Secrets
	smtpClient *lib.SMTPClient
	store      stores.WarrantyStore
}

func NewWarrantyService(store stores.WarrantyStore, smtpClient lib.SMTPClient, secrets *lib.Secrets) *WarrantyServiceImpl {
	return &WarrantyServiceImpl{store: store, smtpClient: &smtpClient, secrets: secrets}
}

func (service *WarrantyServiceImpl) GetWarranties() ([]types.DealerOwnerInfo, error) {
	warranties, err := service.store.GetWarranties()
	if err != nil {
		return nil, err
	}

	return warranties, nil
}

func (service *WarrantyServiceImpl) GetWarrantyById(id string) (*types.WarrantyClaim, []types.PartsRequired, error) {
	warranty, partsRequired, err := service.store.GetWarrantyById(id)
	if err != nil {
		return nil, nil, err
	}

	return warranty, partsRequired, nil
}

func (service *WarrantyServiceImpl) CreateWarranty(warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	if err := service.store.CreateWarranty(warranty, parts); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	if err := service.store.UpdateWarranty(id, warranty, parts); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) DeleteWarranty(id string) error {
	if err := service.store.DeleteWarranty(id); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) SendWarrantyEmail(data *types.EmailData) error {
	return nil
}

func (service *WarrantyServiceImpl) WarrantyFormNotification(client *smtp.Client, data *types.EmailData) error {
	msg := fmt.Sprintf("Subject: New Warranty Form from %s--%s\r\n\r\n%s", data.Name, data.Email, data.Message)

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
