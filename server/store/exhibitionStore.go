package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ExhibitionStore interface {
	GetExhibitions() ([]types.Exhibition, error)
	CreateExhibition(exhibition *types.Exhibition) error
	UpdateExhibition(id string, exhibition *types.Exhibition) error
	DeleteExhibition(id string) error
}

type ExhibitionStoreImpl struct {
	database *sql.DB
}

func NewExhibitionStore(database *sql.DB) *ExhibitionStoreImpl {
	return &ExhibitionStoreImpl{database: database}
}

func (store *ExhibitionStoreImpl) GetExhibitions() ([]types.Exhibition, error) {
	var exhibitions []types.Exhibition

	query := `SELECT * FROM "Exhibition" ORDER BY "created" ASC`
	rows, err := store.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var exhibition types.Exhibition

		err := rows.Scan(&exhibition.ID, &exhibition.Title, &exhibition.Date, &exhibition.Location, &exhibition.Info, &exhibition.Created)
		if err != nil {
			return nil, fmt.Errorf("error occurred while scanning rows: %w", err)
		}

		exhibitions = append(exhibitions, exhibition)
	}

	return exhibitions, nil
}

func (store *ExhibitionStoreImpl) CreateExhibition(exhibition *types.Exhibition) error {
	exhibition.ID = uuid.NewString()
	exhibition.Created = time.Now().String()

	query := `INSERT INTO "Exhibition" (id, title, date, location, info, created) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := store.database.Exec(query, exhibition.ID, exhibition.Title, exhibition.Date, exhibition.Location, exhibition.Info, exhibition.Created)
	if err != nil {
		return fmt.Errorf("error creating exhibition: %w", err)
	}

	return nil
}

func (store *ExhibitionStoreImpl) UpdateExhibition(id string, exhibition *types.Exhibition) error {
	query := `UPDATE "Exhibiton" SET "title" = ?, "date" = ?, "location" = ?, "info" = ? WHERE "id" = ?`

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
