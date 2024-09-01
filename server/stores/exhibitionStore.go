package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
)

type ExhibitionStore interface {
	GetExhibitions(ctx context.Context) ([]db.Exhibition, error)
	CreateExhibition(ctx context.Context, exhibition *db.Exhibition) error
	UpdateExhibition(ctx context.Context, id string, exhibition *db.Exhibition) error
	DeleteExhibition(ctx context.Context, id string) error
}

type ExhibitionStoreImpl struct {
	queries *db.Queries
}

func NewExhibitionStore(sql *sql.DB) *ExhibitionStoreImpl {
	queries := db.New(sql)
	return &ExhibitionStoreImpl{queries: queries}
}

func (store *ExhibitionStoreImpl) GetExhibitions(ctx context.Context) ([]db.Exhibition, error) {
	exhibitions, err := store.queries.GetExhibitions(ctx)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while querying the database for exhibitions: %w", err)
	}

	var result []db.Exhibition
	for _, exhibition := range exhibitions {
		result = append(result, db.Exhibition{
			ID:       exhibition.ID,
			Title:    exhibition.Title,
			Date:     exhibition.Date,
			Location: exhibition.Location,
			Info:     exhibition.Info,
			Created:  exhibition.Created,
		})
	}

	return exhibitions, nil
}

func (store *ExhibitionStoreImpl) CreateExhibition(ctx context.Context, exhibition *db.Exhibition) error {
	exhibition.ID = uuid.NewString()
	exhibition.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}

	params := db.CreateExhibitionParams{
		ID:       exhibition.ID,
		Title:    exhibition.Title,
		Date:     exhibition.Date,
		Location: exhibition.Location,
		Info:     exhibition.Info,
		Created:  exhibition.Created,
	}

	if err := store.queries.CreateExhibition(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating an exhibitions: %w", err)
	}

	return nil
}

func (store *ExhibitionStoreImpl) UpdateExhibition(ctx context.Context, id string, exhibition *db.Exhibition) error {
	params := db.UpdateExhibitionParams{
		Title:    exhibition.Title,
		Date:     exhibition.Date,
		Location: exhibition.Location,
		Info:     exhibition.Info,
		ID:       id,
	}

	if err := store.queries.UpdateExhibition(ctx, params); err != nil {
		return fmt.Errorf("error occurred while updating an exhiibtion: %w", err)
	}
	return nil
}

func (store *ExhibitionStoreImpl) DeleteExhibition(ctx context.Context, id string) error {
	if err := store.queries.DeleteExhibition(ctx, id); err != nil {
		return fmt.Errorf("error occurred while delting an exhibition: %w", err)
	}
	return nil
}
