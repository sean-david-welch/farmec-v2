package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ExhibitionService struct {
	repository *repository.ExhibitionRepository
}

func NewExhibitionService(repository *repository.ExhibitionRepository) *ExhibitionService {
	return &ExhibitionService{repository: repository}
}

func(service *ExhibitionService) GetExhibitions() ([]types.Exhibition, error) {
	exhibitions, err := service.repository.GetExhibitions(); if err != nil {
		return nil, err
	}

	return exhibitions, nil
}

func(service *ExhibitionService) CreateExhibition(exhibition *types.Exhibition) error {
	if err := service.repository.CreateExhibition(exhibition); err != nil {
		return err
	}

	return nil
}

func(service *ExhibitionService) UpdateExhibition(id string, exhibition *types.Exhibition) error {
	if err := service.repository.UpdateExhibition(id, exhibition); err != nil {
		return err
	}

	return nil
}

func(service *ExhibitionService) DeleteExhibition(id string) error {
	if err := service.repository.DeleteExhibition(id); err != nil {
		return err
	}

	return nil
}