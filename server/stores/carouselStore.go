package stores

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type CarouselStore interface {
	GetCarousels() ([]types.Carousel, error)
	GetCarouselById(id string) (*types.Carousel, error)
	CreateCarousel(carousel *types.Carousel) error
	UpdateCarousel(id string, carousel *types.Carousel) error
	DeleteCarousel(id string) error
}

type CarouselStoreImpl struct {
	database *sql.DB
}

func NewCarouselStore(database *sql.DB) *CarouselStoreImpl {
	return &CarouselStoreImpl{database: database}
}

func (store *CarouselStoreImpl) GetCarousels() ([]types.Carousel, error) {
	var carousels []types.Carousel

	query := `SELECT * FROM "Carousel"`
	rows, err := store.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var carousel types.Carousel

		if err := rows.Scan(&carousel.ID, &carousel.Name, &carousel.Image, &carousel.Created); err != nil {
			return nil, fmt.Errorf("error occurred while scanning rows: %w", err)
		}
		carousels = append(carousels, carousel)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return carousels, nil
}

func (store *CarouselStoreImpl) GetCarouselById(id string) (*types.Carousel, error) {
	query := `SELECT * FROM "Carousel" WHERE id = $1`
	row := store.database.QueryRow(query, id)

	var carousel types.Carousel

	if err := row.Scan(&carousel.ID, &carousel.Name, &carousel.Image, &carousel.Created); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &carousel, nil
}

func (store *CarouselStoreImpl) CreateCarousel(carousel *types.Carousel) error {
	carousel.ID = uuid.NewString()
	carousel.Created = time.Now().String()

	query := `INSERT INTO "Carousel"
	(id, name, image, created)
	VALUES ($1, $2, $3, $4)`

	_, err := store.database.Exec(query, carousel.ID, carousel.Name, carousel.Image, carousel.Created)
	if err != nil {
		return fmt.Errorf("error creating carousel: %w", err)
	}

	return nil
}

func (store *CarouselStoreImpl) UpdateCarousel(id string, carousel *types.Carousel) error {
	query := `UPDATE "Carousel" SET name = $1 WHERE id = $3`
	args := []interface{}{carousel.Name, id}

	if carousel.Image != "" && carousel.Image != "null" {
		query = `UPDATE "Carousel" SET name = $1, image = $2 WHERE id = $3`
		args = []interface{}{carousel.Name, carousel.Image, id}
	}

	_, err := store.database.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating carousel: %w", err)
	}

	return nil
}

func (store *CarouselStoreImpl) DeleteCarousel(id string) error {
	query := `DELETE FROM "Carousel" WHERE id = $1`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting supplier: %w", err)
	}

	return nil
}
