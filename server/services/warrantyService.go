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
	GetWarranties(request context.Context) ([]types.DealerOwnerInfo, error)
	GetWarrantyById(request context.Context, id string) (*types.WarrantyClaim, []types.PartsRequired, error)
	CreateWarranty(request context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	UpdateWarranty(request context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	DeleteWarranty(request context.Context, id string) error
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

func (service *WarrantyServiceImpl) GetWarranties(request context.Context) ([]types.DealerOwnerInfo, error) {
	warranties, err := service.repo.GetWarranties(request)
	if err != nil {
		return nil, err
	}

	return warranties, nil
}

func (service *WarrantyServiceImpl) GetWarrantyById(request context.Context, id string) (*types.WarrantyClaim, []types.PartsRequired, error) {
	warranty, partsRequired, err := service.repo.GetWarrantyById(request, id)
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

func (service *WarrantyServiceImpl) CreateWarranty(request context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	if err := service.repo.CreateWarranty(request, warranty, parts); err != nil {
		return err
	}

	go service.sendWarrantyEmail(warranty)

	return nil
}

func (service *WarrantyServiceImpl) UpdateWarranty(request context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	if err := service.repo.UpdateWarranty(request, id, warranty, parts); err != nil {
		return err
	}

	go service.sendWarrantyEmail(warranty)

	return nil
}

func (service *WarrantyServiceImpl) DeleteWarranty(request context.Context, id string) error {
	if err := service.repo.DeleteWarranty(request, id); err != nil {
		return err
	}

	return nil
}
