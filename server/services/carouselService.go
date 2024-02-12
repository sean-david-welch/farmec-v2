package services

import (
	"errors"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type CarouselService interface {
	GetCarousels() ([]types.Carousel, error)
	CreateCarousel(carousel *types.Carousel) (*types.ModelResult, error)
	UpdateCarousel(id string, carousel *types.Carousel) (*types.ModelResult, error)
	DeleteCarousel(id string) error
}

type CarouselServiceImpl struct {
	repository repository.CarouselRepository
	s3Client   utils.S3Client
	folder     string
}

func NewCarouselService(repository repository.CarouselRepository, s3Client utils.S3Client, folder string) *CarouselServiceImpl {
	return &CarouselServiceImpl{
		repository: repository,
		s3Client:   s3Client,
		folder:     folder,
	}
}

func (service *CarouselServiceImpl) GetCarousels() ([]types.Carousel, error) {
	return service.repository.GetCarousels()
}

func (service *CarouselServiceImpl) CreateCarousel(carousel *types.Carousel) (*types.ModelResult, error) {
	image := carousel.Image
	if image == "" || image == "null" {
		return nil, errors.New("image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	carousel.Image = imageUrl

	if err := service.repository.CreateCarousel(carousel); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *CarouselServiceImpl) UpdateCarousel(id string, carousel *types.Carousel) (*types.ModelResult, error) {
	image := carousel.Image

	var presignedUrl, imageUrl string
	var err error

	if image == "" || image == "null" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		carousel.Image = imageUrl
	}

	if err := service.repository.UpdateCarousel(id, carousel); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *CarouselServiceImpl) DeleteCarousel(id string) error {
	carousel, err := service.repository.GetCarouselById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(carousel.Image); err != nil {
		return err
	}

	if err := service.repository.DeleteCarousel(id); err != nil {
		return err
	}

	return nil
}
