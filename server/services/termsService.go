package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsService interface {
	GetTerms() ([]types.Terms, error)
	CreateTerm(term *types.Terms) error
	UpdateTerm(id string, term *types.Terms) error
	DeleteTerm(id string) error
}

type TermsServiceImpl struct {
	repository repository.TermsRepository
}

func NewTermsService(repository repository.TermsRepository) *TermsServiceImpl {
	return &TermsServiceImpl{repository: repository}
}

func(service *TermsServiceImpl) GetTerms() ([]types.Terms, error) {
	terms, err := service.repository.GetTerms(); if err != nil {
		return nil, err
	}
	
	return terms, nil
}

func(service *TermsServiceImpl) CreateTerm(term *types.Terms) error {
	if err := service.repository.CreateTerm(term); err != nil {
		return err
	}
	
	return nil
}

func(service *TermsServiceImpl) UpdateTerm(id string, term *types.Terms) error {
	if err := service.repository.UpdateTerm(id, term); err != nil {
		return err
	}

	return nil
}

func(service *TermsServiceImpl) DeleteTerm(id string) error {
	if err := service.repository.DeleteTerm(id); err != nil {
		return err
	}

	return nil
}