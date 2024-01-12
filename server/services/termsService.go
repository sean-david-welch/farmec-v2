package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
)

type TermService struct {
	repository *repository.TermRepository
}

func NewTermService(repository *repository.TermRepository) *TermService {
	return &TermService{repository: repository}
}

func(service *TermService) GetTerms() ([]models.Terms, error) {
	terms, err := service.repository.GetTerms(); if err != nil {
		return nil, err
	}
	
	return terms, nil
}

func(service *TermService) CreateTerm(term *models.Terms) error {
	if err := service.repository.CreateTerm(term); err != nil {
		return err
	}
	
	return nil
}

func(service *TermService) UpdateTerm(id string, term *models.Terms) error {
	if err := service.repository.UpdateTerm(id, term); err != nil {
		return err
	}

	return nil
}

func(service *TermService) DeleteTerm(id string) error {
	if err := service.repository.DeleteTerm(id); err != nil {
		return err
	}

	return nil
}