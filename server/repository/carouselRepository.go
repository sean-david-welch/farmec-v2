package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type CarouselRepository interface {
	GetCarousels() ([]types.Carousel, error)
	GetCarouselById(id string) (*types.Carousel, error)
	CreateCarousel(carousel *types.Carousel) error
	UpdateCarousel(id string, carousel *types.Carousel) error
	DeleteCarousel(id string) error
}

type CarouselRepositoryImpl struct {
	database *sql.DB
}

func NewCarouselRepository(database *sql.DB) *CarouselRepositoryImpl {
	return &CarouselRepositoryImpl{database: database}
}

func (repository *CarouselRepositoryImpl) GetCarousels() ([]types.Carousel, error) {
	var carousels []types.Carousel

	query := `SELECT * FROM "Carousel"`
	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

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

func (repository *CarouselRepositoryImpl) GetCarouselById(id string) (*types.Carousel, error) {
	query := `SELECT * FROM "Carousel" WHERE id = $1`
	row := repository.database.QueryRow(query, id)

	var carousel types.Carousel

	if err := row.Scan(&carousel.ID, &carousel.Name, &carousel.Image, &carousel.Created); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &carousel, nil
}

func (repository *CarouselRepositoryImpl) CreateCarousel(carousel *types.Carousel) error {
	carousel.ID = uuid.NewString()
	carousel.Created = time.Now()

	query := `INSERT INTO "Carousel"
	(id, name, image, created)
	VALUES ($1, $2, $3, $4)`

	_, err := repository.database.Exec(query, carousel.ID, carousel.Name, carousel.Image, carousel.Created)
	if err != nil {
		return fmt.Errorf("error creating carousel: %w", err)
	}

	return nil
}

func (repository *CarouselRepositoryImpl) UpdateCarousel(id string, carousel *types.Carousel) error {
	query := `UPDATE "Carousel" SET name = $1 WHERE id = $3`
	args := []interface{}{carousel.Name, id}

	if carousel.Image != "" && carousel.Image != "null" {
		query = `UPDATE "Carousel" SET name = $1, image = $2 WHERE id = $3`
		args = []interface{}{carousel.Name, carousel.Image, id}
	}

	_, err := repository.database.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating carousel: %w", err)
	}

	return nil
}

func (repository *CarouselRepositoryImpl) DeleteCarousel(id string) error {
	query := `DELETE FROM "Carousel" WHERE id = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting supplier: %w", err)
	}

	return nil
}
