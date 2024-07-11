package services

import (
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"
	url2 "net/url"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PartsService interface {
	GetParts(id string) ([]types.Sparepart, error)
	CreatePart(part *types.Sparepart) (*types.PartsModelResult, error)
	UpdatePart(id string, part *types.Sparepart) (*types.PartsModelResult, error)
	DeletePart(id string) error
}

type PartsServiceImpl struct {
	folder   string
	s3Client lib.S3Client
	store    stores.PartsStore
}

func NewPartsService(store stores.PartsStore, s3Client lib.S3Client, folder string) *PartsServiceImpl {
	return &PartsServiceImpl{
		store:    store,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *PartsServiceImpl) GetParts(id string) ([]types.Sparepart, error) {
	return service.store.GetParts(id)
}

func (service *PartsServiceImpl) CreatePart(part *types.Sparepart) (*types.PartsModelResult, error) {
	partsImage := part.PartsImage
	if partsImage == "" || partsImage == "null" {
		return nil, errors.New("image is empty")
	}
	presignedImageUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, partsImage)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}
	part.PartsImage = imageUrl

	var partsLinkUrl, presignedLinkUrl string
	partsLink := part.SparePartsLink
	u, err := url2.Parse(partsLink)
	if err == nil && u.Scheme != "" && u.Host != "" {
		partsLinkUrl = partsLink
	} else {
		presignedUrl, linkUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, partsLink)
		if err != nil {
			log.Printf("error occurred while generating presigned url for link: %v", err)
			return nil, err
		}
		partsLinkUrl = linkUrl
		presignedLinkUrl = presignedUrl
	}
	part.SparePartsLink = partsLinkUrl

	err = service.store.CreatePart(part)
	if err != nil {
		log.Printf("Failed to create part: %v", err)
		return nil, err
	}

	result := &types.PartsModelResult{
		PresignedUrl:     presignedImageUrl,
		ImageUrl:         imageUrl,
		PresignedLinkUrl: presignedLinkUrl,
	}

	return result, nil
}

func (service *PartsServiceImpl) UpdatePart(id string, part *types.Sparepart) (*types.PartsModelResult, error) {
	partsImage := part.PartsImage
	var presignedUrl, imageUrl string
	var err error

	if partsImage != "" && partsImage != "null" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, partsImage)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		part.PartsImage = imageUrl
	}

	var partsLinkUrl, presignedLinkUrl string
	partsLink := part.SparePartsLink

	u, err := url2.Parse(partsLink)
	if err == nil && u.Scheme != "" && u.Host != "" {
		partsLinkUrl = partsLink
	} else {
		presignedUrl, linkUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, partsLink)
		if err != nil {
			log.Printf("error occurred while generating presigned url for link: %v", err)
			return nil, err
		}
		partsLinkUrl = linkUrl
		presignedLinkUrl = presignedUrl
	}
	part.SparePartsLink = partsLinkUrl

	err = service.store.UpdatePart(id, part)
	if err != nil {
		log.Printf("Failed to update part: %v", err)
		return nil, err
	}

	result := &types.PartsModelResult{
		PresignedUrl:     presignedUrl,
		ImageUrl:         imageUrl,
		PresignedLinkUrl: presignedLinkUrl,
	}

	return result, nil
}

func (service *PartsServiceImpl) DeletePart(id string) error {
	part, err := service.store.GetPartById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(part.PartsImage); err != nil {
		return err
	}

	if err := service.store.DeletePart(id); err != nil {
		return err
	}

	return nil
}
