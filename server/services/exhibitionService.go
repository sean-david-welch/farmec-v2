package services

import (
	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ExhibitionService interface {
	GetExhibitions() ([]types.Exhibition, error)
	CreateExhibition(exhibition *types.Exhibition) error
	UpdateExhibition(id string, exhibition *types.Exhibition) error
	DeleteExhibition(id string) error
}

type ExhibitionServiceImpl struct {
	store store.ExhibitionStore
}

func NewExhibitionService(store store.ExhibitionStore) *ExhibitionServiceImpl {
	return &ExhibitionServiceImpl{store: store}
}

func (service *ExhibitionServiceImpl) GetExhibitions() ([]types.Exhibition, error) {
	exhibitions, err := service.store.GetExhibitions()
	if err != nil {
		return nil, err
	}

	return exhibitions, nil
}

func (service *ExhibitionServiceImpl) CreateExhibition(exhibition *types.Exhibition) error {
	if err := service.store.CreateExhibition(exhibition); err != nil {
		return err
	}

	return nil
}

func (service *ExhibitionServiceImpl) UpdateExhibition(id string, exhibition *types.Exhibition) error {
	if err := service.store.UpdateExhibition(id, exhibition); err != nil {
		return err
	}

	return nil
}

func (service *ExhibitionServiceImpl) DeleteExhibition(id string) error {
	if err := service.store.DeleteExhibition(id); err != nil {
		return err
	}

	return nil
}
