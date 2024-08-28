package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsStore interface {
	GetTerms() ([]types.Terms, error)
	CreateTerm(term *types.Terms) error
	UpdateTerm(id string, term *types.Terms) error
	DeleteTerm(id string) error
}

type TermsStoreImpl struct {
	database *sql.DB
}

func NewTermsStore(database *sql.DB) *TermsStoreImpl {
	return &TermsStoreImpl{database: database}
}

func (store *TermsStoreImpl) GetTerms() ([]types.Terms, error) {
	var terms []types.Terms

	query := `SELECT * FROM "Terms"`
	rows, err := store.database.Query(query)
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

func (store *TermsStoreImpl) CreateTerm(term *types.Terms) error {
	term.ID = uuid.NewString()
	term.Created = time.Now().String()

	query := `INSERT INTO "Terms" (id, title, body, created) VALUES (?, ?, ?, ?)`

	_, err := store.database.Exec(query, term.ID, term.Title, term.Body, term.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating term term: %w", err)
	}

	return nil
}

func (store *TermsStoreImpl) UpdateTerm(id string, term *types.Terms) error {
	query := `UPDATE "Terms" SET title = ?, body = ? where id = ?`

	_, err := store.database.Exec(query, term.Title, term.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating term term: %w", err)
	}

	return nil
}

func (store *TermsStoreImpl) DeleteTerm(id string) error {
	query := `DELETE FROM "Terms" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting term term: %w", err)
	}

	return nil
}
