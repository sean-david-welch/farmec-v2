package services

import (
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type CarouselService interface {
	GetCarousels() ([]types.Carousel, error)
	CreateCarousel(carousel *types.Carousel) (*types.ModelResult, error)
	UpdateCarousel(id string, carousel *types.Carousel) (*types.ModelResult, error)
	DeleteCarousel(id string) error
}

type CarouselServiceImpl struct {
	store    stores.CarouselStore
	s3Client lib.S3Client
	folder   string
}

func NewCarouselService(store stores.CarouselStore, s3Client lib.S3Client, folder string) *CarouselServiceImpl {
	return &CarouselServiceImpl{
		store:    store,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *CarouselServiceImpl) GetCarousels() ([]types.Carousel, error) {
	return service.store.GetCarousels()
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

	if err := service.store.CreateCarousel(carousel); err != nil {
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

	if image != "" && image != "null" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		carousel.Image = imageUrl
	}

	if err := service.store.UpdateCarousel(id, carousel); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *CarouselServiceImpl) DeleteCarousel(id string) error {
	carousel, err := service.store.GetCarouselById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(carousel.Image); err != nil {
		return err
	}

	if err := service.store.DeleteCarousel(id); err != nil {
		return err
	}

	return nil
}
