package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type PrivacyRepo interface {
	GetPrivacy(ctx context.Context) ([]db.Privacy, error)
	CreatePrivacy(ctx context.Context, privacy *db.Privacy) error
	UpdatePrivacy(ctx context.Context, id string, privacy *db.Privacy) error
	DeletePrivacy(ctx context.Context, id string) error
}

type PrivacyRepoImpl struct {
	queries *db.Queries
}

func NewPrivacyRepo(sql *sql.DB) *PrivacyRepoImpl {
	queries := db.New(sql)
	return &PrivacyRepoImpl{queries: queries}
}

func (repo *PrivacyRepoImpl) GetPrivacy(ctx context.Context) ([]db.Privacy, error) {
	privacies, err := repo.queries.GetPrivacies(ctx)
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

func (repo *PrivacyRepoImpl) CreatePrivacy(ctx context.Context, privacy *db.Privacy) error {
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
	if err := repo.queries.CreatePrivacy(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating policy: %w", err)
	}
	return nil
}

func (repo *PrivacyRepoImpl) UpdatePrivacy(ctx context.Context, id string, privacy *db.Privacy) error {
	params := db.UpdatePrivacyParams{
		Title: privacy.Title,
		Body:  privacy.Body,
		ID:    id,
	}
	if err := repo.queries.UpdatePrivacy(ctx, params); err != nil {
		return fmt.Errorf("an error occurred while")
	}
	return nil
}

func (repo *PrivacyRepoImpl) DeletePrivacy(ctx context.Context, id string) error {
	if err := repo.queries.DeletePrivacy(ctx, id); err != nil {
		return fmt.Errorf("an error occurred while deleting privacy: %w", err)
	}
	return nil
}
