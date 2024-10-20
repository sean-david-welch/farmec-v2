package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PrivacyService interface {
	GetPrivacys(ctx context.Context) ([]types.Privacy, error)
	CreatePrivacy(ctx context.Context, privacy *db.Privacy) error
	UpdatePrivacy(ctx context.Context, id string, privacy *db.Privacy) error
	DeletePrivacy(ctx context.Context, id string) error
}

type PrivacyServiceImpl struct {
	repo repository.PrivacyRepo
}

func NewPrivacyService(repo repository.PrivacyRepo) *PrivacyServiceImpl {
	return &PrivacyServiceImpl{repo: repo}
}

func (service *PrivacyServiceImpl) GetPrivacys(ctx context.Context) ([]types.Privacy, error) {
	privacys, err := service.repo.GetPrivacy(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.Privacy
	for _, privacy := range privacys {
		result = append(result, lib.SerializePrivacy(privacy))
	}
	return result, nil
}

func (service *PrivacyServiceImpl) CreatePrivacy(ctx context.Context, privacy *db.Privacy) error {
	if err := service.repo.CreatePrivacy(ctx, privacy); err != nil {
		return err
	}

	return nil
}

func (service *PrivacyServiceImpl) UpdatePrivacy(ctx context.Context, id string, privacy *db.Privacy) error {
	if err := service.repo.UpdatePrivacy(ctx, id, privacy); err != nil {
		return err
	}

	return nil
}

func (service *PrivacyServiceImpl) DeletePrivacy(ctx context.Context, id string) error {
	if err := service.repo.DeletePrivacy(ctx, id); err != nil {
		return err
	}

	return nil
}
