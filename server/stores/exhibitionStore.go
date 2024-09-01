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
		return nil, fmt.Errorf("An error occurred while querying the database for machines: %w", err)
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
}

func (store *ExhibitionStoreImpl) CreateExhibition(exhibition *db.Exhibition) error {
	exhibition.ID = uuid.NewString()
	exhibition.Created = time.Now().String()

	query := `INSERT INTO "Exhibition" (id, title, date, location, info, created) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := store.database.Exec(query, exhibition.ID, exhibition.Title, exhibition.Date, exhibition.Location, exhibition.Info, exhibition.Created)
	if err != nil {
		return fmt.Errorf("error creating exhibition: %w", err)
	}

	return nil
}

func (store *ExhibitionStoreImpl) UpdateExhibition(id string, exhibition *db.Exhibition) error {
	query := `UPDATE "Exhibition" SET "title" = ?, "date" = ?, "location" = ?, "info" = ? WHERE "id" = ?`

	_, err := store.database.Exec(query, id, exhibition.Title, exhibition.Date, exhibition.Location, exhibition.Info)
	if err != nil {
		return fmt.Errorf("error updating exhibition: %w", err)
	}

	return nil
}

func (store *ExhibitionStoreImpl) DeleteExhibition(id string) error {
	query := `DELETE FROM "Exhibition" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting exhibiton: %w", err)
	}

	return nil
}
