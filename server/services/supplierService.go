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

type SupplierService interface {
	GetSuppliers(ctx context.Context) ([]types.Supplier, error)
	CreateSupplier(ctx context.Context, supplier *db.Supplier) (*types.SupplierResult, error)
	GetSupplierById(ctx context.Context, id string) (*types.Supplier, error)
	UpdateSupplier(ctx context.Context, id string, supplier *db.Supplier) (*types.SupplierResult, error)
	DeleteSupplier(ctx context.Context, id string) error
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

func (service *SupplierServiceImpl) GetSuppliers(ctx context.Context) ([]types.Supplier, error) {
	suppliers, err := service.store.GetSuppliers(ctx)
	if err != nil {
		return nil, err
	}

	var result []types.Supplier
	for _, supplier := range suppliers {
		result = append(result, types.Supplier{
			ID:              supplier.ID,
			Name:            supplier.Name,
			LogoImage:       supplier.LogoImage.String,
			MarketingImage:  supplier.MarketingImage.String,
			Description:     supplier.Description.String,
			SocialFacebook:  &supplier.SocialFacebook.String,
			SocialInstagram: &supplier.SocialInstagram.String,
			SocialLinkedin:  &supplier.SocialLinkedin.String,
			SocialTwitter:   &supplier.SocialTwitter.String,
			SocialYoutube:   &supplier.SocialYoutube.String,
			SocialWebsite:   &supplier.SocialWebsite.String,
			Created:         supplier.Created.String,
		})
	}
	return result, nil
}

func (service *SupplierServiceImpl) GetSupplierById(ctx context.Context, id string) (*types.Supplier, error) {
	supplier, err := service.store.GetSupplierById(ctx, id)
	if err != nil {
		return nil, err
	}
	result := &types.Supplier{
		ID:              supplier.ID,
		Name:            supplier.Name,
		LogoImage:       supplier.LogoImage.String,
		MarketingImage:  supplier.MarketingImage.String,
		Description:     supplier.Description.String,
		SocialFacebook:  &supplier.SocialFacebook.String,
		SocialInstagram: &supplier.SocialInstagram.String,
		SocialLinkedin:  &supplier.SocialLinkedin.String,
		SocialTwitter:   &supplier.SocialTwitter.String,
		SocialYoutube:   &supplier.SocialYoutube.String,
		SocialWebsite:   &supplier.SocialWebsite.String,
		Created:         supplier.Created.String,
	}
	return result, nil
}

func (service *SupplierServiceImpl) CreateSupplier(ctx context.Context, supplier *db.Supplier) (*types.SupplierResult, error) {
	logoImage := supplier.LogoImage
	marketingImage := supplier.MarketingImage

	if !logoImage.Valid || !marketingImage.Valid {
		return nil, errors.New("image is empty")
	}

	presignedLogo, logoUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, logoImage.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	presignedMarketing, marketingUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, marketingImage.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	supplier.LogoImage = sql.NullString{
		String: logoUrl,
		Valid:  true,
	}
	supplier.MarketingImage = sql.NullString{
		String: marketingUrl,
		Valid:  true,
	}

	err = service.store.CreateSupplier(ctx, supplier)
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

func (service *SupplierServiceImpl) UpdateSupplier(ctx context.Context, id string, supplier *db.Supplier) (*types.SupplierResult, error) {
	logoImage := supplier.LogoImage
	marketingImage := supplier.MarketingImage

	var presignedLogo, logoUrl, presignedMarketing, marketingUrl string
	var err error

	if logoImage.Valid {
		presignedLogo, logoUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, logoImage.String)
		if err != nil {
			return nil, err
		}
		supplier.LogoImage = sql.NullString{
			String: logoUrl,
			Valid:  true,
		}
	}

	if marketingImage.Valid {
		presignedMarketing, marketingUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, marketingImage.String)
		if err != nil {
			return nil, err
		}
		supplier.MarketingImage = sql.NullString{
			String: marketingUrl,
			Valid:  true,
		}
	}

	err = service.store.UpdateSupplier(ctx, id, supplier)
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

func (service *SupplierServiceImpl) DeleteSupplier(ctx context.Context, id string) error {
	supplier, err := service.store.GetSupplierById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(supplier.LogoImage.String); err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(supplier.MarketingImage.String); err != nil {
		return err
	}

	if err := service.store.DeleteSupplier(ctx, id); err != nil {
		return err
	}

	return nil
}
