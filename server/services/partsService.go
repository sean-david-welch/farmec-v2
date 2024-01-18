package services

import (
	"errors"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type PartsService interface {
	GetParts(id string) ([]types.Sparepart, error) 
	CreatePart(part *types.Sparepart) (*types.ModelResult, error) 
	UpdatePart(id string, part *types.Sparepart) (*types.ModelResult, error) 
	DeletePart(id string) error 
}

type PartsServiceImpl struct {
	folder string
	s3Client *utils.S3Client
	repository repository.PartsRepository
}

func NewPartsService(repository repository.PartsRepository, s3Client *utils.S3Client, folder string) *PartsServiceImpl {
	return &PartsServiceImpl{
		repository: repository,
		s3Client: s3Client,
		folder: folder,
	}
}

func (service *PartsServiceImpl) GetParts(id string) ([]types.Sparepart, error) {
	return service.repository.GetParts(id)
}

func (service *PartsServiceImpl) CreatePart(part *types.Sparepart) (*types.ModelResult, error) {
	partsImage := part.PartsImage; if partsImage == "" {
		return nil, errors.New("parts image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, partsImage)
	if err != nil {
		return nil, err
	}

	part.PartsImage =  imageUrl

	service.repository.CreatePart(part); if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: imageUrl,
	}

	return result, nil
}

func (service *PartsServiceImpl) UpdatePart(id string, part *types.Sparepart) (*types.ModelResult, error) {
	partsImage := part.PartsImage

	var presignedUrl, imageUrl string
	var err error

	if partsImage != "" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, partsImage)
		if err != nil {
			return nil, err
		}
		part.PartsImage = imageUrl
	}

	service.repository.UpdatePart(id, part); if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: imageUrl,
	}

	return result, nil
}

func (service *PartsServiceImpl) DeletePart(id string) error {
	part, err := service.repository.GetPartById(id); if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(part.PartsImage); err != nil {
		return err
	}
	
	if err := service.repository.DeletePart(id); err != nil {
		return err
	}

	return nil
}