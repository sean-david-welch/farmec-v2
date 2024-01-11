package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

// type Carousel struct {
//     ID    string `json:"id"`
//     Name  string `json:"name"`
//     Image string `json:"image"`
// }

type CarouselRepository struct {
	db *sql.DB
}

func NewCarouselRepository(db *sql.DB) *CarouselRepository {
	return &CarouselRepository{db: db}
}

func (repository *CarouselRepository) GetCarousels() ([]models.Carousel, error) {
	var carousels []models.Carousel

	query := `SELECT * FROM "Carousel"`
	rows, err := repository.db.Query(query); if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var carousel models.Carousel

		if err := rows.Scan(&carousel.ID, &carousel.Name, &carousel.Image); err != nil {
			return nil, fmt.Errorf("")
		}
		carousels = append(carousels, carousel)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return carousels, nil
}

func (repository *CarouselRepository) GetCarouselById(id string) (*models.Carousel, error) {
	query := `SELECT * FROM "Carousel" WHERE id = $1`
	row := repository.db.QueryRow(query, id)

	var carousel models.Carousel

	if err := row.Scan(&carousel.ID, &carousel.Name, &carousel.Image); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &carousel, nil
}

func (repository *CarouselRepository) CreateCarousel(carousel *models.Carousel) error {
	carousel.ID = uuid.NewString()

	query := `INSERT INTO "Carousel"
	(id, name, image)
	VALUES ($1, $2, $3)`

	_, err := repository.db.Exec(query, &carousel.ID, &carousel.Name, &carousel.Image); if err != nil {
		return fmt.Errorf("error creating carousel: %w", err)
	}
	
	return nil
}

func (repository *CarouselRepository) UpdateCarousel(id string, carousel *models.Carousel) error {
	query := `UPDATE "Carousel" SET name = $1, image = $2 WHERE id = $3`

	_, err := repository.db.Exec(query, &carousel.Name, &carousel.Image, id); if err != nil {
		return fmt.Errorf("error updating carousel: %w", err)
	}

	return nil
}

func (repository *CarouselRepository) DeleteCarousel(id string) error {
	query := `DELETE FROM "Carousel" WHERE id = $1`

	_, err := repository.db.Exec(query, id); if err != nil {
		return fmt.Errorf("error deleting supplier: %w", err)
	}
	
	return nil
}