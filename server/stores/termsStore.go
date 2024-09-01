package stores

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type TermsStore interface {
	GetTerms(ctx context.Context) ([]db.Term, error)
	CreateTerm(ctx context.Context, term *db.Term) error
	UpdateTerm(ctx context.Context, id string, term *db.Term) error
	DeleteTerm(ctx context.Context, id string) error
}

type TermsStoreImpl struct {
	queries *db.Queries
}

func NewTermsStore(sql *sql.DB) *TermsStoreImpl {
	queries := db.New(sql)
	return &TermsStoreImpl{queries: queries}
}

func (store *TermsStoreImpl) GetTerms(ctx context.Context) ([]db.Term, error) {
	terms, err := store.queries.GetTerms(ctx)
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

func (store *TermsStoreImpl) CreateTerm(ctx context.Context, term *db.Term) error {
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
	if err := store.queries.CreateTerm(ctx, params); err != nil {
		return fmt.Errorf("error occured while creating a term: %w", err)
	}
	return nil
}

func (store *TermsStoreImpl) UpdateTerm(id string, term *db.Term) error {
	query := `UPDATE "Term" SET title = ?, body = ? where id = ?`

	_, err := store.database.Exec(query, term.Title, term.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating term term: %w", err)
	}

	return nil
}

func (store *TermsStoreImpl) DeleteTerm(id string) error {
	query := `DELETE FROM "Term" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting term term: %w", err)
	}

	return nil
}
