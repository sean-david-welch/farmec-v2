package services

import (
	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsService interface {
	GetTerms() ([]types.Terms, error)
	CreateTerm(term *types.Terms) error
	UpdateTerm(id string, term *types.Terms) error
	DeleteTerm(id string) error
}

type TermsServiceImpl struct {
	store stores.TermsStore
}

func NewTermsService(store stores.TermsStore) *TermsServiceImpl {
	return &TermsServiceImpl{store: store}
}

func (service *TermsServiceImpl) GetTerms() ([]types.Terms, error) {
	terms, err := service.store.GetTerms()
	if err != nil {
		return nil, err
	}

	return terms, nil
}

func (service *TermsServiceImpl) CreateTerm(term *types.Terms) error {
	if err := service.store.CreateTerm(term); err != nil {
		return err
	}

	return nil
}

func (service *TermsServiceImpl) UpdateTerm(id string, term *types.Terms) error {
	if err := service.store.UpdateTerm(id, term); err != nil {
		return err
	}

	return nil
}

func (service *TermsServiceImpl) DeleteTerm(id string) error {
	if err := service.store.DeleteTerm(id); err != nil {
		return err
	}

	return nil
}
