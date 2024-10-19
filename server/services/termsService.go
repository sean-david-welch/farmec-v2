package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsService interface {
	GetTerms(ctx context.Context) ([]types.Terms, error)
	CreateTerm(ctx context.Context, term *db.Term) error
	UpdateTerm(ctx context.Context, id string, term *db.Term) error
	DeleteTerm(ctx context.Context, id string) error
}

type TermsServiceImpl struct {
	store repository.TermsRepo
}

func NewTermsService(store repository.TermsRepo) *TermsServiceImpl {
	return &TermsServiceImpl{store: store}
}

func (service *TermsServiceImpl) GetTerms(ctx context.Context) ([]types.Terms, error) {
	terms, err := service.store.GetTerms(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.Terms
	for _, term := range terms {
		result = append(result, lib.SerializeTerm(term))
	}
	return result, nil
}

func (service *TermsServiceImpl) CreateTerm(ctx context.Context, term *db.Term) error {
	if err := service.store.CreateTerm(ctx, term); err != nil {
		return err
	}
	return nil
}

func (service *TermsServiceImpl) UpdateTerm(ctx context.Context, id string, term *db.Term) error {
	if err := service.store.UpdateTerm(ctx, id, term); err != nil {
		return err
	}
	return nil
}

func (service *TermsServiceImpl) DeleteTerm(ctx context.Context, id string) error {
	if err := service.store.DeleteTerm(ctx, id); err != nil {
		return err
	}
	return nil
}
