package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
)

type WarrantyService interface {
	GetWarranties(reqContext context.Context) ([]types.DealerOwnerInfo, error)
	GetWarrantyById(reqContext context.Context, id string) (*types.WarrantyClaim, []types.PartsRequired, error)
	CreateWarranty(reqContext context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	UpdateWarranty(reqContext context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	DeleteWarranty(reqContext context.Context, id string) error
}

type WarrantyServiceImpl struct {
	repo        repository.WarrantyRepo
	emailClient *lib.EmailClientImpl
}

func NewWarrantyService(repo repository.WarrantyRepo, emailClient *lib.EmailClientImpl) *WarrantyServiceImpl {
	return &WarrantyServiceImpl{repo: repo, emailClient: emailClient}
}

func (service *WarrantyServiceImpl) sendWarrantyEmail(warranty *db.WarrantyClaim) {
	data := &types.EmailData{
		Name:    warranty.OwnerName,
		Email:   warranty.Dealer,
		Message: warranty.MachineModel,
	}

	if err := service.emailClient.SendFormNotification(data, "Warranty"); err != nil {
		log.Printf("Failed to send warranty email: %v", err)
		return
	}
}

func (service *WarrantyServiceImpl) GetWarranties(reqContext context.Context) ([]types.DealerOwnerInfo, error) {
	warranties, err := service.repo.GetWarranties(reqContext)
	if err != nil {
		return nil, err
	}

	return warranties, nil
}

func (service *WarrantyServiceImpl) GetWarrantyById(reqContext context.Context, id string) (*types.WarrantyClaim, []types.PartsRequired, error) {
	warranty, partsRequired, err := service.repo.GetWarrantyById(reqContext, id)
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

func (service *WarrantyServiceImpl) CreateWarranty(reqContext context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	if err := service.repo.CreateWarranty(reqContext, warranty, parts); err != nil {
		return err
	}

	go service.sendWarrantyEmail(warranty)

	return nil
}

func (service *WarrantyServiceImpl) UpdateWarranty(reqContext context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	if err := service.repo.UpdateWarranty(reqContext, id, warranty, parts); err != nil {
		return err
	}

	go service.sendWarrantyEmail(warranty)

	return nil
}

func (service *WarrantyServiceImpl) DeleteWarranty(reqContext context.Context, id string) error {
	if err := service.repo.DeleteWarranty(reqContext, id); err != nil {
		return err
	}

	return nil
}
