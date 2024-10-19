package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net/smtp"
)

type WarrantyService interface {
	GetWarranties(ctx context.Context) ([]types.DealerOwnerInfo, error)
	GetWarrantyById(ctx context.Context, id string) (*types.WarrantyClaim, []types.PartsRequired, error)
	CreateWarranty(ctx context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	UpdateWarranty(ctx context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	DeleteWarranty(ctx context.Context, id string) error
}

type WarrantyServiceImpl struct {
	smtpClient lib.SMTPClient
	store      repository.WarrantyStore
}

func NewWarrantyService(store repository.WarrantyStore, smtpClient lib.SMTPClient) *WarrantyServiceImpl {
	return &WarrantyServiceImpl{store: store, smtpClient: smtpClient}
}

func (service *WarrantyServiceImpl) GetWarranties(ctx context.Context) ([]types.DealerOwnerInfo, error) {
	warranties, err := service.store.GetWarranties(ctx)
	if err != nil {
		return nil, err
	}

	return warranties, nil
}

func (service *WarrantyServiceImpl) GetWarrantyById(ctx context.Context, id string) (*types.WarrantyClaim, []types.PartsRequired, error) {
	warranty, partsRequired, err := service.store.GetWarrantyById(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	warrantyClaim := lib.SerializeWarrantyClaim(*warranty)
	var result []types.PartsRequired
	for _, part := range partsRequired {
		result = append(result, lib.SerializePartsRequired(part))
	}
	return &warrantyClaim, result, nil
}

func (service *WarrantyServiceImpl) CreateWarranty(ctx context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	if err := service.store.CreateWarranty(ctx, warranty, parts); err != nil {
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
		Name:    warranty.OwnerName,
		Email:   warranty.Dealer,
		Message: warranty.MachineModel,
	}

	if err := service.smtpClient.SendFormNotification(client, data, "Warranty"); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) UpdateWarranty(ctx context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	if err := service.store.UpdateWarranty(ctx, id, warranty, parts); err != nil {
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
		Name:    warranty.OwnerName,
		Email:   warranty.Dealer,
		Message: warranty.MachineModel,
	}

	if err := service.smtpClient.SendFormNotification(client, data, "Warranty"); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) DeleteWarranty(ctx context.Context, id string) error {
	if err := service.store.DeleteWarranty(ctx, id); err != nil {
		return err
	}

	return nil
}
