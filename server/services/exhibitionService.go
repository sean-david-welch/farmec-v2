package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ExhibitionService interface {
	GetExhibitions() ([]types.Exhibition, error)
	CreateExhibition(exhibition *types.Exhibition) error
	UpdateExhibition(id string, exhibition *types.Exhibition) error
	DeleteExhibition(id string) error
}

type ExhibitionServiceImpl struct {
	repository repository.ExhibitionRepository
}

func NewExhibitionService(repository repository.ExhibitionRepository) *ExhibitionServiceImpl {
	return &ExhibitionServiceImpl{repository: repository}
}

func(service *ExhibitionServiceImpl) GetExhibitions() ([]types.Exhibition, error) {
	exhibitions, err := service.repository.GetExhibitions(); if err != nil {
		return nil, err
	}

	return exhibitions, nil
}

func(service *ExhibitionServiceImpl) CreateExhibition(exhibition *types.Exhibition) error {
	if err := service.repository.CreateExhibition(exhibition); err != nil {
		return err
	}

	return nil
}

func(service *ExhibitionServiceImpl) UpdateExhibition(id string, exhibition *types.Exhibition) error {
	if err := service.repository.UpdateExhibition(id, exhibition); err != nil {
		return err
	}

	return nil
}

func(service *ExhibitionServiceImpl) DeleteExhibition(id string) error {
	if err := service.repository.DeleteExhibition(id); err != nil {
		return err
	}

	return nil
}