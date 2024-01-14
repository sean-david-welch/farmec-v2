package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

type PrivacyRepository struct {
	database *sql.DB
}

// type Privacy struct {
//     ID      string `json:"id"`
//     Title   string `json:"title"`
//     Body    string `json:"body"`
//     Created time.Time `json:"created"`
// }

// type Terms struct {
//     ID      string `json:"id"`
//     Title   string `json:"title"`
//     Body    string `json:"body"`
//     Created time.Time `json:"created"`
// }

func NewPrivacyRepository(database *sql.DB) *PrivacyRepository {
	return &PrivacyRepository{database: database}
}

func(repository *PrivacyRepository) GetPrivacy() ([]models.Privacy, error) {
	var privacyTerms []models.Privacy

	query := `SELECT * FROM "Privacy"`
	rows, err := repository.database.Query(query); if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var privacyTerm models.Privacy

		if err := rows.Scan(&privacyTerm.ID, privacyTerm.Title, privacyTerm.Body, privacyTerm.Created); err != nil {
			return nil, fmt.Errorf("error occurred while scanning row: %w", err)
		}

		privacyTerms = append(privacyTerms, privacyTerm)

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error occurred after iterating over rows: %w", err)
		}	
	}

	return privacyTerms, err
}

func(repository *PrivacyRepository) CreatePrivacy(privacy *models.Privacy) error {
	privacy.ID = uuid.NewString()
	privacy.Created = time.Now()

	query := `INSERT INTO "Privacy" (id, title, body, created) VALUES ($1, $2, $3, $4)`

	_, err := repository.database.Exec(query, privacy.ID, privacy.Title, privacy.Body, privacy.Created); if err != nil {
		return fmt.Errorf("error occurred while creating privacy term: %w", err)
	}
	
	return nil
}

func(repository *PrivacyRepository) UpdatePrivacy(id string, privacy *models.Privacy) error {
	query := `UPDATE "Privacy" SET title = $1, body = $2 where id = $3`

	_, err := repository.database.Exec(query, privacy.Title, privacy.Body, id); if err != nil {
		return fmt.Errorf("error occurred while updating privacy term: %w", err)
	}

	return nil
}

func(repository *PrivacyRepository) DeletePrivacy(id string) error {
	query := `DELETE FROM "Privacy" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id); if err != nil {
		return fmt.Errorf("error occurred while deleting privacy term: %w", err)
	}

	return nil
}