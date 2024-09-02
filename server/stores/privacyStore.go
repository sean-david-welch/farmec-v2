package stores

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type PrivacyStore interface {
	GetPrivacy(ctx context.Context) ([]db.Privacy, error)
	CreatePrivacy(ctx context.Context, privacy *db.Privacy) error
	UpdatePrivacy(ctx context.Context, id string, privacy *db.Privacy) error
	DeletePrivacy(ctx context.Context, id string) error
}

type PrivacyStoreImpl struct {
	queries *db.Queries
}

func NewPrivacyStore(sql *sql.DB) *PrivacyStoreImpl {
	queries := db.New(sql)
	return &PrivacyStoreImpl{queries: queries}
}

func (store *PrivacyStoreImpl) GetPrivacy(ctx context.Context) ([]db.Privacy, error) {
	privacies, err := store.queries.GetPrivacies(ctx)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting privacy policy: %w", err)
	}
	var result []db.Privacy
	for _, privacy := range privacies {
		result = append(result, db.Privacy{
			ID:      privacy.ID,
			Title:   privacy.Title,
			Body:    privacy.Body,
			Created: privacy.Created,
		})
	}
	return result, nil
}

func (store *PrivacyStoreImpl) CreatePrivacy(ctx context.Context, privacy *db.Privacy) error {
	privacy.ID = uuid.NewString()
	privacy.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}

	params := db.CreatePrivacyParams{
		ID:      privacy.ID,
		Title:   privacy.Title,
		Body:    privacy.Body,
		Created: privacy.Created,
	}
	if err := store.queries.CreatePrivacy(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating policy: %w", err)
	}
	return nil
}

func (store *PrivacyStoreImpl) UpdatePrivacy(ctx context.Context, id string, privacy *db.Privacy) error {
	query := `UPDATE "Privacy" SET title = ?, body = ? where id = ?`

	_, err := store.database.Exec(query, privacy.Title, privacy.Body, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating privacy term: %w", err)
	}

	return nil
}

func (store *PrivacyStoreImpl) DeletePrivacy(ctx context.Context, id string) error {
	query := `DELETE FROM "Privacy" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting privacy term: %w", err)
	}

	return nil
}
