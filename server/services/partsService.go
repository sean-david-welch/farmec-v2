package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"
	url2 "net/url"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PartsService interface {
	GetParts(ctx context.Context, id string) ([]types.Sparepart, error)
	GetPartsSupplier(ctx context.Context, id string) (*types.Supplier, error)
	CreatePart(ctx context.Context, part *db.SparePart) (*types.PartsModelResult, error)
	UpdatePart(ctx context.Context, id string, part *db.SparePart) (*types.PartsModelResult, error)
	DeletePart(ctx context.Context, id string) error
}

type PartsServiceImpl struct {
	folder   string
	s3Client lib.S3Client
	repo     repository.PartsRepo
}

func NewPartsService(repo repository.PartsRepo, s3Client lib.S3Client, folder string) *PartsServiceImpl {
	return &PartsServiceImpl{
		repo:     repo,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *PartsServiceImpl) GetParts(ctx context.Context, id string) ([]types.Sparepart, error) {
	parts, err := service.repo.GetParts(ctx, id)
	if err != nil {
		return nil, err
	}
	var result []types.Sparepart
	for _, part := range parts {
		result = append(result, lib.SerializeSparePart(part))
	}
	return result, nil
}

func (service *PartsServiceImpl) GetPartsSupplier(ctx context.Context, id string) (*types.Supplier, error) {
	supplier, err := service.repo.GetPartsSupplier(ctx, id)
	if err != nil {
		return nil, err
	}
	result := lib.SerializeSupplier(*supplier)
	return &result, nil
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

	err = service.repo.CreatePart(ctx, part)
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

	err = service.repo.UpdatePart(ctx, id, part)
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
	part, err := service.repo.GetPartById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(part.PartsImage.String); err != nil {
		return err
	}

	if err := service.repo.DeletePart(ctx, id); err != nil {
		return err
	}

	return nil
}
