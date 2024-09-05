package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/stores"
)

type ExhibitionService interface {
	GetExhibitions(ctx context.Context) ([]db.Exhibition, error)
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

func (service *ExhibitionServiceImpl) GetExhibitions(ctx context.Context) ([]db.Exhibition, error) {
	exhibitions, err := service.store.GetExhibitions(ctx)
	if err != nil {
		return nil, err
	}

	return exhibitions, nil
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
