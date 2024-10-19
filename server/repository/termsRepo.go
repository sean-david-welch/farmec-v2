package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type TermsRepo interface {
	GetTerms(ctx context.Context) ([]db.Term, error)
	CreateTerm(ctx context.Context, term *db.Term) error
	UpdateTerm(ctx context.Context, id string, term *db.Term) error
	DeleteTerm(ctx context.Context, id string) error
}

type TermsRepoImpl struct {
	queries *db.Queries
}

func NewTermsRepo(sql *sql.DB) *TermsRepoImpl {
	queries := db.New(sql)
	return &TermsRepoImpl{queries: queries}
}

func (repo *TermsRepoImpl) GetTerms(ctx context.Context) ([]db.Term, error) {
	terms, err := repo.queries.GetTerms(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying the database for terms: %w", err)
	}
	var result []db.Term
	for _, term := range terms {
		result = append(result, db.Term{
			ID:      term.ID,
			Title:   term.Title,
			Body:    term.Body,
			Created: term.Created,
		})
	}
	return result, nil
}

func (repo *TermsRepoImpl) CreateTerm(ctx context.Context, term *db.Term) error {
	term.ID = uuid.NewString()
	term.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}
	params := db.CreateTermParams{
		ID:      term.ID,
		Title:   term.Title,
		Body:    term.Body,
		Created: term.Created,
	}
	if err := repo.queries.CreateTerm(ctx, params); err != nil {
		return fmt.Errorf("error occured while creating a term: %w", err)
	}
	return nil
}

func (repo *TermsRepoImpl) UpdateTerm(ctx context.Context, id string, term *db.Term) error {
	params := db.UpdateTermParams{
		Title:   term.Title,
		Body:    term.Body,
		Created: term.Created,
		ID:      id,
	}
	if err := repo.queries.UpdateTerm(ctx, params); err != nil {
		return fmt.Errorf("error ocurred while updating a machine with image: %w", err)
	}
	return nil
}

func (repo *TermsRepoImpl) DeleteTerm(ctx context.Context, id string) error {
	if err := repo.queries.DeleteTerm(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting term: %w", err)
	}
	return nil
}
