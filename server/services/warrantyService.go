package services

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
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
	smtpClient lib.SMTPClient
	store      store.WarrantyStore
}

func NewWarrantyService(store store.WarrantyStore, smtpClient lib.SMTPClient) *WarrantyServiceImpl {
	return &WarrantyServiceImpl{store: store, smtpClient: smtpClient}
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

	data := &types.EmailData{
		Name:    *warranty.OwnerName,
		Email:   warranty.Dealer,
		Message: *warranty.MachineModel,
	}

	if err := service.smtpClient.SendFormNotification(client, data, "Warranty"); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	if err := service.store.UpdateWarranty(id, warranty, parts); err != nil {
		return err
	}

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

	data := &types.EmailData{
		Name:    *warranty.OwnerName,
		Email:   warranty.Dealer,
		Message: *warranty.MachineModel,
	}

	if err := service.smtpClient.SendFormNotification(client, data, "Warranty"); err != nil {
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
