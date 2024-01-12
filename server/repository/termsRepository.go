package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

type TermsRepository struct {
	db *sql.DB
}

func NewTermsRepository(db *sql.DB) *TermsRepository {
	return &TermsRepository{db: db}
}

func(repository *TermsRepository) GetTerms() ([]models.Terms, error) {
	var terms []models.Terms

	query := `SELECT * FROM "Terms"`
	rows, err := repository.db.Query(query); if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var term models.Terms

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

func(repository *TermsRepository) CreateTerm(term *models.Terms) error {
	term.ID = uuid.NewString()
	term.Created = time.Now()

	query := `INSERT INTO "Terms" (id, title, body, created) VALUES ($1, $2, $3, $4)`

	_, err := repository.db.Exec(query, &term.ID, &term.Title, &term.Body, &term.Created); if err != nil {
		return fmt.Errorf("error occurred while creating term term: %w", err)
	}
	
	return nil
}

func(repository *TermsRepository) UpdateTerm(id string, term *models.Terms) error {
	query := `UPDATE "Terms" SET title = $1, body = $2 where id = $3`

	_, err := repository.db.Exec(query, &term.Title, &term.Body, id); if err != nil {
		return fmt.Errorf("error occurred while updating term term: %w", err)
	}

	return nil
}

func(repository *TermsRepository) DeleteTerm(id string) error {
	query := `DELETE FROM "Terms" WHERE "id" = $1`

	_, err := repository.db.Exec(query, id); if err != nil {
		return fmt.Errorf("error occurred while deleting term term: %w", err)
	}

	return nil
}