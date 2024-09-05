package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"
	url2 "net/url"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PartsService interface {
	GetParts(ctx context.Context, id string) ([]db.SparePart, error)
	CreatePart(ctx context.Context, part *db.SparePart) (*types.PartsModelResult, error)
	UpdatePart(ctx context.Context, id string, part *db.SparePart) (*types.PartsModelResult, error)
	DeletePart(ctx context.Context, id string) error
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

func (service *PartsServiceImpl) GetParts(ctx context.Context, id string) ([]db.SparePart, error) {
	parts, err := service.store.GetParts(ctx, id)
	if err != nil {
		return nil, err
	}
	return parts, nil
}

func (service *PartsServiceImpl) CreatePart(ctx context.Context, part *db.SparePart) (*types.PartsModelResult, error) {
	partsImage := part.PartsImage
	if !partsImage.Valid {
		return nil, errors.New("image is empty")
	}

	presignedImageUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, partsImage.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}
	part.PartsImage = sql.NullString{
		String: imageUrl,
		Valid:  true,
	}

	var partsLinkUrl, presignedLinkUrl string
	partsLink := sql.NullString{
		String: part.SparePartsLink.String,
		Valid:  true,
	}
	u, err := url2.Parse(partsLink.String)
	if err == nil && u.Scheme != "" && u.Host != "" {
		partsLinkUrl = partsLink.String
	} else {
		presignedUrl, linkUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, partsLink.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url for link: %v", err)
			return nil, err
		}
		partsLinkUrl = linkUrl
		presignedLinkUrl = presignedUrl
	}
	part.SparePartsLink = sql.NullString{
		String: partsLinkUrl,
		Valid:  true,
	}

	err = service.store.CreatePart(ctx, part)
	if err != nil {
		log.Printf("Failed to create part: %v", err)
		return nil, err
	}

	result := &types.PartsModelResult{
		PresignedImageUrl: presignedImageUrl,
		ImageUrl:          imageUrl,
		PresignedLinkUrl:  presignedLinkUrl,
	}

	return result, nil
}

func (service *PartsServiceImpl) UpdatePart(ctx context.Context, id string, part *db.SparePart) (*types.PartsModelResult, error) {
	partsImage := part.PartsImage
	var presignedUrl, imageUrl string
	var err error

	if partsImage.Valid {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, partsImage.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		part.PartsImage = sql.NullString{
			String: imageUrl,
			Valid:  true,
		}
	}

	var partsLinkUrl, presignedLinkUrl string
	partsLink := sql.NullString{
		String: part.SparePartsLink.String,
		Valid:  true,
	}

	u, err := url2.Parse(partsLink.String)
	if err == nil && u.Scheme != "" && u.Host != "" {
		partsLinkUrl = partsLink.String
	} else {
		presignedUrl, linkUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, partsLink.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url for link: %v", err)
			return nil, err
		}
		partsLinkUrl = linkUrl
		presignedLinkUrl = presignedUrl
	}
	part.SparePartsLink = sql.NullString{
		String: partsLinkUrl,
		Valid:  true,
	}

	err = service.store.UpdatePart(ctx, id, part)
	if err != nil {
		log.Printf("Failed to update part: %v", err)
		return nil, err
	}

	result := &types.PartsModelResult{
		PresignedImageUrl: presignedUrl,
		ImageUrl:          imageUrl,
		PresignedLinkUrl:  presignedLinkUrl,
	}

	return result, nil
}

func (service *PartsServiceImpl) DeletePart(ctx context.Context, id string) error {
	part, err := service.store.GetPartById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(part.PartsImage.String); err != nil {
		return err
	}

	if err := service.store.DeletePart(ctx, id); err != nil {
		return err
	}

	return nil
}
