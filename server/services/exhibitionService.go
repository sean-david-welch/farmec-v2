package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ExhibitionService interface {
	GetExhibitions(ctx context.Context) ([]types.Exhibition, error)
	CreateExhibition(ctx context.Context, exhibition *db.Exhibition) error
	UpdateExhibition(ctx context.Context, id string, exhibition *db.Exhibition) error
	DeleteExhibition(ctx context.Context, id string) error
}

type ExhibitionServiceImpl struct {
	store stores.ExhibitionStore
}

func NewExhibitionService(store stores.ExhibitionStore) *ExhibitionServiceImpl {
	return &ExhibitionServiceImpl{store: store}
}

func (service *ExhibitionServiceImpl) GetExhibitions(ctx context.Context) ([]types.Exhibition, error) {
	exhibitions, err := service.store.GetExhibitions(ctx)
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
	if err := service.store.CreateExhibition(ctx, exhibition); err != nil {
		return err
	}

	return nil
}

func (service *ExhibitionServiceImpl) UpdateExhibition(ctx context.Context, id string, exhibition *db.Exhibition) error {
	if err := service.store.UpdateExhibition(ctx, id, exhibition); err != nil {
		return err
	}

	return nil
}

func (service *ExhibitionServiceImpl) DeleteExhibition(ctx context.Context, id string) error {
	if err := service.store.DeleteExhibition(ctx, id); err != nil {
		return err
	}

	return nil
}
