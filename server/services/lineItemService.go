package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemService interface {
	GetLineItems(ctx context.Context) ([]types.LineItem, error)
	GetLineItemById(ctx context.Context, id string) (*types.LineItem, error)
	CreateLineItem(ctx context.Context, lineItem *db.LineItem) (*types.ModelResult, error)
	UpdateLineItem(ctx context.Context, id string, lineItem *db.LineItem) (*types.ModelResult, error)
	DeleteLineItem(ctx context.Context, id string) error
}

type LineItemServiceImpl struct {
	store    stores.LineItemStore
	s3Client lib.S3Client
	folder   string
}

func NewLineItemService(store stores.LineItemStore, s3Client lib.S3Client, folder string) *LineItemServiceImpl {
	return &LineItemServiceImpl{store: store, s3Client: s3Client, folder: folder}
}

func (service *LineItemServiceImpl) GetLineItems(ctx context.Context) ([]types.LineItem, error) {
	lineItems, err := service.store.GetLineItems(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.LineItem
	for _, lineItem := range lineItems {
		result = append(result, lib.SerializeLineItem(lineItem))
	}
	return result, nil
}

func (service *LineItemServiceImpl) GetLineItemById(ctx context.Context, id string) (*types.LineItem, error) {
	lineItem, err := service.store.GetLineItemById(ctx, id)
	if err != nil {
		return nil, err
	}
	result := lib.SerializeLineItem(*lineItem)
	return &result, nil
}

func (service *LineItemServiceImpl) CreateLineItem(ctx context.Context, lineItem *db.LineItem) (*types.ModelResult, error) {
	image := lineItem.Image

	if !image.Valid {
		return nil, errors.New("image is empty")
	}

	PresignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	lineItem.Image = sql.NullString{
		String: imageUrl,
		Valid:  true,
	}

	if err := service.store.CreateLineItem(ctx, lineItem); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *LineItemServiceImpl) UpdateLineItem(ctx context.Context, id string, lineItem *db.LineItem) (*types.ModelResult, error) {
	image := lineItem.Image

	var PresignedUrl, imageUrl string
	var err error

	if image.Valid {
		PresignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		lineItem.Image = sql.NullString{
			String: imageUrl,
			Valid:  true,
		}
	}

	if err := service.store.UpdateLineItem(ctx, id, lineItem); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *LineItemServiceImpl) DeleteLineItem(ctx context.Context, id string) error {
	lineItem, err := service.store.GetLineItemById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(lineItem.Image.String); err != nil {
		return err
	}

	if err := service.store.DeleteLineItem(ctx, id); err != nil {
		return err
	}

	return nil
}
