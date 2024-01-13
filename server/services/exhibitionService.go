package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
)

type ExhibitionService struct {
	repository *repository.ExhibitionRepository
}

func NewExhibitionService(repository *repository.ExhibitionRepository) *ExhibitionService {
	return &ExhibitionService{repository: repository}
}

func(service *ExhibitionService) GetExhibitions() ([]models.Exhibition, error) {
	exhibitions, err := service.repository.GetExhibitions(); if err != nil {
		return nil, err
	}

	return exhibitions, nil
}

func(service *ExhibitionService) CreateExhibition(exhibition *models.Exhibition) error {
	if err := service.repository.CreateExhibition(exhibition); err != nil {
		return err
	}

	return nil
}

func(service *ExhibitionService) UpdateExhibition(id string, exhibition *models.Exhibition) error {
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