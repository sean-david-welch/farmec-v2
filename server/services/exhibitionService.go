package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ExhibitionService interface {
	GetExhibitions(ctx context.Context) ([]types.Exhibition, error)
	CreateExhibition(ctx context.Context, exhibition *db.Exhibition) error
	UpdateExhibition(ctx context.Context, id string, exhibition *db.Exhibition) error
	DeleteExhibition(ctx context.Context, id string) error
}

type ExhibitionServiceImpl struct {
	repo repository.ExhibitionRepo
}

func NewExhibitionService(repo repository.ExhibitionRepo) *ExhibitionServiceImpl {
	return &ExhibitionServiceImpl{repo: repo}
}

func (service *ExhibitionServiceImpl) GetExhibitions(ctx context.Context) ([]types.Exhibition, error) {
	exhibitions, err := service.repo.GetExhibitions(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.Exhibition
	for _, exhibition := range exhibitions {
		result = append(result, lib.SerializeExhibition(exhibition))
	}
	return result, nil
}

func (service *ExhibitionServiceImpl) CreateExhibition(ctx context.Context, exhibition *db.Exhibition) error {
	if err := service.repo.CreateExhibition(ctx, exhibition); err != nil {
		return err
	}

	return nil
}

func (service *ExhibitionServiceImpl) UpdateExhibition(ctx context.Context, id string, exhibition *db.Exhibition) error {
	if err := service.repo.UpdateExhibition(ctx, id, exhibition); err != nil {
		return err
	}

	return nil
}

func (service *ExhibitionServiceImpl) DeleteExhibition(ctx context.Context, id string) error {
	if err := service.repo.DeleteExhibition(ctx, id); err != nil {
		return err
	}

	return nil
}
