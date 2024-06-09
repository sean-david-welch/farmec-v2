package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsRepository interface {
	GetTerms() ([]types.Terms, error)
	CreateTerm(term *types.Terms) error
	UpdateTerm(id string, term *types.Terms) error
	DeleteTerm(id string) error
}

type TermsRepositoryImpl struct {
	database *sql.DB
}

func NewTermsRepository(database *sql.DB) *TermsRepositoryImpl {
	return &TermsRepositoryImpl{database: database}
}

func (repository *TermsRepositoryImpl) GetTerms() ([]types.Terms, error) {
	var terms []types.Terms

	query := `SELECT * FROM "Terms"`
	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var term types.Terms

		if err := rows.Scan(&term.ID, &term.Title, &term.Body, &term.Created); err != nil {
			return nil, fmt.Errorf("error occurred while scanning row: %w", err)
		}

		terms = append(terms, term)

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error occurred after iterating over rows: %w", err)
		}
	}

	return terms, err
}

func (repository *TermsRepositoryImpl) CreateTerm(term *types.Terms) error {
	term.ID = uuid.NewString()
	term.Created = time.Now()

	query := `INSERT INTO "Terms" (id, title, body, created) VALUES ($1, $2, $3, $4)`

	_, err := repository.database.Exec(query, term.ID, term.Title, term.Body, term.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating term term: %w", err)
	}

	return nil
}

func (repository *TermsRepositoryImpl) UpdateTerm(id string, term *types.Terms) error {
	query := `UPDATE "Terms" SET title = $1, body = $2 where id = $3`

	_, err := repository.database.Exec(query, term.Title, term.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating term term: %w", err)
	}

	return nil
}

func (repository *TermsRepositoryImpl) DeleteTerm(id string) error {
	query := `DELETE FROM "Terms" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting term term: %w", err)
	}

	return nil
}
