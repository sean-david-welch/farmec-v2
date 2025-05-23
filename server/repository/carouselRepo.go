package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
)

type CarouselRepo interface {
	GetCarousels(ctx context.Context) ([]db.Carousel, error)
	GetCarouselById(ctx context.Context, id string) (*db.Carousel, error)
	CreateCarousel(ctx context.Context, carousel *db.Carousel) error
	UpdateCarousel(ctx context.Context, id string, carousel *db.Carousel) error
	DeleteCarousel(ctx context.Context, id string) error
}

type CarouselRepoImpl struct {
	queries *db.Queries
}

func NewCarouselRepo(sql *sql.DB) *CarouselRepoImpl {
	queries := db.New(sql)
	return &CarouselRepoImpl{queries: queries}
}

func (repo *CarouselRepoImpl) GetCarousels(ctx context.Context) ([]db.Carousel, error) {
	carousels, err := repo.queries.GetCarousels(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying the database for carousels: %w", err)
	}

	var result []db.Carousel
	for _, carousel := range carousels {
		result = append(result, db.Carousel{
			ID:      carousel.ID,
			Name:    carousel.Name,
			Image:   carousel.Image,
			Created: carousel.Created,
		})
	}
	return result, nil
}

func (repo *CarouselRepoImpl) GetCarouselById(ctx context.Context, id string) (*db.Carousel, error) {
	carousel, err := repo.queries.GetCarouselByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying the database for carousel: %w", err)
	}

	return &carousel, nil
}

func (repo *CarouselRepoImpl) CreateCarousel(ctx context.Context, carousel *db.Carousel) error {
	carousel.ID = uuid.NewString()
	carousel.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}

	params := db.CreateCarouselParams{
		ID:      carousel.ID,
		Name:    carousel.Name,
		Image:   carousel.Image,
		Created: carousel.Created,
	}

	if err := repo.queries.CreateCarousel(ctx, params); err != nil {
		return fmt.Errorf("error occurred while create a carousel: %w", err)
	}
	return nil
}

func (repo *CarouselRepoImpl) UpdateCarousel(ctx context.Context, id string, carousel *db.Carousel) error {
	if carousel.Image.Valid {
		params := db.UpdateCarouselParams{
			Name:  carousel.Name,
			Image: carousel.Image,
			ID:    id,
		}
		if err := repo.queries.UpdateCarousel(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updating a carousel: %w", err)
		}
	} else {
		params := db.UpdateCarouselNoImageParams{
			Name: carousel.Name,
			ID:   id,
		}
		if err := repo.queries.UpdateCarouselNoImage(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updating a carousel: %w", err)
		}
	}
	return nil
}

func (repo *CarouselRepoImpl) DeleteCarousel(ctx context.Context, id string) error {
	if err := repo.queries.DeleteCarousel(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting a carousel: %w", err)
	}
	return nil
}
