package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsService struct {
	repository *repository.TermsRepository
}

func NewTermsService(repository *repository.TermsRepository) *TermsService {
	return &TermsService{repository: repository}
}

func(service *TermsService) GetTerms() ([]types.Terms, error) {
	terms, err := service.repository.GetTerms(); if err != nil {
		return nil, err
	}
	
	return terms, nil
}

func(service *TermsService) CreateTerm(term *types.Terms) error {
	if err := service.repository.CreateTerm(term); err != nil {
		return err
	}
	
	return nil
}

func(service *TermsService) UpdateTerm(id string, term *types.Terms) error {
	if err := service.repository.UpdateTerm(id, term); err != nil {
		return err
	}

	return nil
}

func(service *TermsService) DeleteTerm(id string) error {
	if err := service.repository.DeleteTerm(id); err != nil {
		return err
	}

	return nil
}