package services

import (
	"errors"

	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type CarouselService struct {
	repository *repository.CarouselRepository
	s3Client *utils.S3Client
	folder string
}

func NewCarouselService(repository *repository.CarouselRepository, s3Client *utils.S3Client, folder string) *CarouselService {
	return &CarouselService{
		repository: repository,
		s3Client: s3Client,
		folder: folder,
	}
}

func (service *CarouselService) GetCarousels() ([]models.Carousel, error) {
	return service.repository.GetCarousels()
}

func (service *CarouselService) CreateCarousel(carousel *models.Carousel) (*types.ModelResult, error) {
	image := carousel.Image; if image != "" {
		return nil, errors.New("image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image)
	if err != nil {
		return nil, err
	}

	carousel.Image = imageUrl

	service.repository.CreateCarousel(carousel); if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: imageUrl,
	}

	return result, nil
}	

func (service *CarouselService) UpdateCarousel(id string, carousel *models.Carousel) (*types.ModelResult, error) {
	image := carousel.Image;

	var presignedUrl, imageUrl string
	var err error

	if image != "" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image)
		if err != nil {
			return nil, err
		}
		carousel.Image = imageUrl
	}

	service.repository.UpdateCarousel(id, carousel); if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: imageUrl,
	}

	return result, nil
}

func (service *CarouselService) DeleteCarousel(id string) error {
	carousel, err := service.repository.GetCarouselById(id); if err != nil {
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

