package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PrivacyRepository interface {
	GetPrivacy() ([]types.Privacy, error)
	CreatePrivacy(privacy *types.Privacy) error
	UpdatePrivacy(id string, privacy *types.Privacy) error
	DeletePrivacy(id string) error
}

type PrivacyRepositoryImpl struct {
	database *sql.DB
}

func NewPrivacyRepository(database *sql.DB) *PrivacyRepositoryImpl {
	return &PrivacyRepositoryImpl{database: database}
}

func (repository *PrivacyRepositoryImpl) GetPrivacy() ([]types.Privacy, error) {
	var privacyTerms []types.Privacy

	query := `SELECT * FROM "Privacy"`
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
		var privacyTerm types.Privacy

		if err := rows.Scan(&privacyTerm.ID, &privacyTerm.Title, &privacyTerm.Body, &privacyTerm.Created); err != nil {
			return nil, fmt.Errorf("error occurred while scanning row: %w", err)
		}

		privacyTerms = append(privacyTerms, privacyTerm)

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error occurred after iterating over rows: %w", err)
		}
	}

	return privacyTerms, err
}

func (repository *PrivacyRepositoryImpl) CreatePrivacy(privacy *types.Privacy) error {
	privacy.ID = uuid.NewString()
	privacy.Created = time.Now()

	query := `INSERT INTO "Privacy" (id, title, body, created) VALUES ($1, $2, $3, $4)`

	_, err := repository.database.Exec(query, privacy.ID, privacy.Title, privacy.Body, privacy.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating privacy term: %w", err)
	}

	return nil
}

func (repository *PrivacyRepositoryImpl) UpdatePrivacy(id string, privacy *types.Privacy) error {
	query := `UPDATE "Privacy" SET title = $1, body = $2 where id = $3`

	_, err := repository.database.Exec(query, privacy.Title, privacy.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating privacy term: %w", err)
	}

	return nil
}

func (repository *PrivacyRepositoryImpl) DeletePrivacy(id string) error {
	query := `DELETE FROM "Privacy" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting privacy term: %w", err)
	}

	return nil
}
