package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PrivacyStore interface {
	GetPrivacy() ([]types.Privacy, error)
	CreatePrivacy(privacy *types.Privacy) error
	UpdatePrivacy(id string, privacy *types.Privacy) error
	DeletePrivacy(id string) error
}

type PrivacyStoreImpl struct {
	database *sql.DB
}

func NewPrivacyStore(database *sql.DB) *PrivacyStoreImpl {
	return &PrivacyStoreImpl{database: database}
}

func (store *PrivacyStoreImpl) GetPrivacy() ([]types.Privacy, error) {
	var privacyTerms []types.Privacy

	query := `SELECT * FROM "Privacy"`
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

func (store *PrivacyStoreImpl) CreatePrivacy(privacy *types.Privacy) error {
	privacy.ID = uuid.NewString()
	privacy.Created = time.Now().String()

	query := `INSERT INTO "Privacy" (id, title, body, created) VALUES (?, ?, ?, ?)`

	_, err := store.database.Exec(query, privacy.ID, privacy.Title, privacy.Body, privacy.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating privacy term: %w", err)
	}

	return nil
}

func (store *PrivacyStoreImpl) UpdatePrivacy(id string, privacy *types.Privacy) error {
	query := `UPDATE "Privacy" SET title = ?, body = ? where id = ?`

	_, err := store.database.Exec(query, privacy.Title, privacy.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating privacy term: %w", err)
	}

	return nil
}

func (store *PrivacyStoreImpl) DeletePrivacy(id string) error {
	query := `DELETE FROM "Privacy" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting privacy term: %w", err)
	}

	return nil
}
