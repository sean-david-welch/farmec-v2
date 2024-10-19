package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/repository"
)

type CarouselService interface {
	GetCarousels(ctx context.Context) ([]types.Carousel, error)
	CreateCarousel(ctx context.Context, carousel *db.Carousel) (*types.ModelResult, error)
	UpdateCarousel(ctx context.Context, id string, carousel *db.Carousel) (*types.ModelResult, error)
	DeleteCarousel(ctx context.Context, id string) error
}

type CarouselServiceImpl struct {
	repo     repository.CarouselRepo
	s3Client lib.S3Client
	folder   string
}

func NewCarouselService(repo repository.CarouselRepo, s3Client lib.S3Client, folder string) *CarouselServiceImpl {
	return &CarouselServiceImpl{
		repo:     repo,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *CarouselServiceImpl) GetCarousels(ctx context.Context) ([]types.Carousel, error) {
	carousels, err := service.repo.GetCarousels(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.Carousel
	for _, carousel := range carousels {
		result = append(result, lib.SerializeCarousel(carousel))
	}
	return result, nil
}

func (service *CarouselServiceImpl) CreateCarousel(ctx context.Context, carousel *db.Carousel) (*types.ModelResult, error) {
	image := carousel.Image

	if !image.Valid {
		return nil, errors.New("image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	carousel.Image = sql.NullString{
		String: imageUrl,
		Valid:  true,
	}

	if err := service.repo.CreateCarousel(ctx, carousel); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *CarouselServiceImpl) UpdateCarousel(ctx context.Context, id string, carousel *db.Carousel) (*types.ModelResult, error) {
	image := carousel.Image

	var presignedUrl, imageUrl string
	var err error

	if image.Valid {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		carousel.Image = sql.NullString{
			String: imageUrl,
			Valid:  true,
		}
	}

	if err := service.repo.UpdateCarousel(ctx, id, carousel); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *CarouselServiceImpl) DeleteCarousel(ctx context.Context, id string) error {
	carousel, err := service.repo.GetCarouselById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(carousel.Image.String); err != nil {
		return err
	}

	if err := service.repo.DeleteCarousel(ctx, id); err != nil {
		return err
	}

	return nil
}
