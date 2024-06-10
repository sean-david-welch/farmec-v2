package services

import (
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type SupplierService interface {
	GetSuppliers() ([]types.Supplier, error)
	CreateSupplier(supplier *types.Supplier) (*types.SupplierResult, error)
	GetSupplierById(id string) (*types.Supplier, error)
	UpdateSupplier(id string, supplier *types.Supplier) (*types.SupplierResult, error)
	DeleteSupplier(id string) error
}

type SupplierServiceImpl struct {
	folder   string
	s3Client lib.S3Client
	store    stores.SupplierStore
}

func NewSupplierService(store stores.SupplierStore, s3Client lib.S3Client, folder string) *SupplierServiceImpl {
	return &SupplierServiceImpl{
		store:    store,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *SupplierServiceImpl) GetSuppliers() ([]types.Supplier, error) {
	return service.store.GetSuppliers()
}

func (service *SupplierServiceImpl) CreateSupplier(supplier *types.Supplier) (*types.SupplierResult, error) {
	logoImage := supplier.LogoImage
	marketingImage := supplier.MarketingImage

	if logoImage == "" || logoImage == "null" {
		return nil, errors.New("image is empty")
	}

	presignedLogo, logoUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, logoImage)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	presignedMarketing, marketingUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, marketingImage)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	supplier.LogoImage = logoUrl
	supplier.MarketingImage = marketingUrl

	err = service.store.CreateSupplier(supplier)
	if err != nil {
		return nil, err
	}

	result := &types.SupplierResult{
		PresignedLogoUrl:      presignedLogo,
		LogoUrl:               logoUrl,
		PresignedMarketingUrl: presignedMarketing,
		MarketingUrl:          marketingUrl,
	}

	return result, nil
}

func (service *SupplierServiceImpl) GetSupplierById(id string) (*types.Supplier, error) {
	return service.store.GetSupplierById(id)
}

func (service *SupplierServiceImpl) UpdateSupplier(id string, supplier *types.Supplier) (*types.SupplierResult, error) {
	logoImage := supplier.LogoImage
	marketingImage := supplier.MarketingImage

	var presignedLogo, logoUrl, presignedMarketing, marketingUrl string
	var err error

	if logoImage != "" && logoImage != "null" {
		presignedLogo, logoUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, logoImage)
		if err != nil {
			return nil, err
		}
		supplier.LogoImage = logoUrl
	}

	if marketingImage != "" && marketingImage != "null" {
		presignedMarketing, marketingUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, marketingImage)
		if err != nil {
			return nil, err
		}
		supplier.MarketingImage = marketingUrl
	}

	err = service.store.UpdateSupplier(id, supplier)
	if err != nil {
		return nil, err
	}

	result := &types.SupplierResult{
		PresignedLogoUrl:      presignedLogo,
		LogoUrl:               logoUrl,
		PresignedMarketingUrl: presignedMarketing,
		MarketingUrl:          marketingUrl,
	}

	return result, nil
}

func (service *SupplierServiceImpl) DeleteSupplier(id string) error {
	supplier, err := service.store.GetSupplierById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(supplier.LogoImage); err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(supplier.MarketingImage); err != nil {
		return err
	}

	if err := service.store.DeleteSupplier(id); err != nil {
		return err
	}

	return nil
}
