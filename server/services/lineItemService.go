package services

import (
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemService interface {
	GetLineItems() ([]types.LineItem, error)
	GetLineItemById(id string) (*types.LineItem, error)
	CreateLineItem(lineItem *types.LineItem) (*types.ModelResult, error)
	UpdateLineItem(id string, lineItem *types.LineItem) (*types.ModelResult, error)
	DeleteLineItem(id string) error
}

type LineItemServiceImpl struct {
	store    stores.LineItemStore
	s3Client lib.S3Client
	folder   string
}

func NewLineItemService(store stores.LineItemStore, s3Client lib.S3Client, folder string) *LineItemServiceImpl {
	return &LineItemServiceImpl{store: store, s3Client: s3Client, folder: folder}
}

func (service *LineItemServiceImpl) GetLineItems() ([]types.LineItem, error) {
	lineItems, err := service.store.GetLineItems()
	if err != nil {
		return nil, err
	}

	return lineItems, nil
}

func (service *LineItemServiceImpl) GetLineItemById(id string) (*types.LineItem, error) {
	lineItem, err := service.store.GetLineItemById(id)
	if err != nil {
		return nil, err
	}

	return lineItem, nil
}

func (service *LineItemServiceImpl) CreateLineItem(lineItem *types.LineItem) (*types.ModelResult, error) {
	image := lineItem.Image

	if image == "" || image == "null" {
		return nil, errors.New("image is empty")
	}

	PresignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	lineItem.Image = imageUrl

	if err := service.store.CreateLineItem(lineItem); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *LineItemServiceImpl) UpdateLineItem(id string, lineItem *types.LineItem) (*types.ModelResult, error) {
	image := lineItem.Image

	var PresignedUrl, imageUrl string
	var err error

	if image != "" && image != "null" {
		PresignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		lineItem.Image = imageUrl
	}

	if err := service.store.UpdateLineItem(id, lineItem); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *LineItemServiceImpl) DeleteLineItem(id string) error {
	lineItem, err := service.store.GetLineItemById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(lineItem.Image); err != nil {
		return err
	}

	if err := service.store.DeleteLineItem(id); err != nil {
		return err
	}

	return nil
}
