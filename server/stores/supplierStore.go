package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
)

type SupplierStore interface {
	GetSuppliers(ctx context.Context) ([]db.Supplier, error)
	CreateSupplier(ctx context.Context, supplier *db.Supplier) error
	GetSupplierById(ctx context.Context, id string) (*db.Supplier, error)
	UpdateSupplier(ctx context.Context, id string, supplier *db.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error
}

type SupplierStoreImpl struct {
	queries *db.Queries
}

func NewSupplierStore(sql *sql.DB) *SupplierStoreImpl {
	queries := db.New(sql)
	return &SupplierStoreImpl{queries: queries}
}

func (store *SupplierStoreImpl) GetSuppliers(ctx context.Context) ([]db.Supplier, error) {
	suppliers, err := store.queries.GetSuppliers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting suppliers from the db: %w", err)
	}

	var result []db.Supplier
	for _, supplier := range suppliers {
		result = append(result, db.Supplier{
			ID:              supplier.ID,
			Name:            supplier.Name,
			LogoImage:       supplier.LogoImage,
			MarketingImage:  supplier.MarketingImage,
			Description:     supplier.Description,
			SocialFacebook:  supplier.SocialFacebook,
			SocialTwitter:   supplier.SocialTwitter,
			SocialInstagram: supplier.SocialInstagram,
			SocialYoutube:   supplier.SocialYoutube,
			SocialLinkedin:  supplier.SocialLinkedin,
			SocialWebsite:   supplier.SocialWebsite,
			Created:         supplier.Created,
		})
	}
	return result, nil
}

func (store *SupplierStoreImpl) GetSupplierById(ctx context.Context, id string) (*db.Supplier, error) {
	supplierRow, err := store.queries.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting suppliers from the db: %w", err)
	}
	supplier := db.Supplier{
		ID:              supplierRow.ID,
		Name:            supplierRow.Name,
		LogoImage:       supplierRow.LogoImage,
		MarketingImage:  supplierRow.MarketingImage,
		Description:     supplierRow.Description,
		SocialFacebook:  supplierRow.SocialFacebook,
		SocialTwitter:   supplierRow.SocialTwitter,
		SocialInstagram: supplierRow.SocialInstagram,
		SocialYoutube:   supplierRow.SocialYoutube,
		SocialLinkedin:  supplierRow.SocialLinkedin,
		SocialWebsite:   supplierRow.SocialWebsite,
		Created:         supplierRow.Created,
	}
	return &supplier, nil
}

func (store *SupplierStoreImpl) CreateSupplier(ctx context.Context, supplier *db.Supplier) error {
	supplier.ID = uuid.NewString()
	supplier.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}

	params := db.CreateSupplierParams{
		ID:              supplier.ID,
		Name:            supplier.Name,
		LogoImage:       supplier.LogoImage,
		MarketingImage:  supplier.MarketingImage,
		Description:     supplier.Description,
		SocialFacebook:  supplier.SocialFacebook,
		SocialTwitter:   supplier.SocialTwitter,
		SocialInstagram: supplier.SocialInstagram,
		SocialYoutube:   supplier.SocialYoutube,
		SocialLinkedin:  supplier.SocialLinkedin,
		SocialWebsite:   supplier.SocialWebsite,
		Created:         supplier.Created,
	}
	if err := store.queries.CreateSupplier(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating supplier: %w", err)
	}
	return nil
}

func (store *SupplierStoreImpl) UpdateSupplier(ctx context.Context, id string, supplier *db.Supplier) error {
	if supplier.LogoImage.Valid && supplier.MarketingImage.Valid {
		params := db.UpdateSupplierParams{
			Name:            supplier.Name,
			LogoImage:       supplier.LogoImage,
			MarketingImage:  supplier.MarketingImage,
			Description:     supplier.Description,
			SocialFacebook:  supplier.SocialFacebook,
			SocialTwitter:   supplier.SocialTwitter,
			SocialInstagram: supplier.SocialInstagram,
			SocialYoutube:   supplier.SocialYoutube,
			SocialLinkedin:  supplier.SocialLinkedin,
			SocialWebsite:   supplier.SocialWebsite,
			ID:              id,
		}
		if err := store.queries.UpdateSupplier(ctx, params); err != nil {
			return fmt.Errorf("error while updating the supplier: %w", err)
		}
	} else {
		params := db.UpdateSupplierNoImageParams{
			Name:            supplier.Name,
			Description:     supplier.Description,
			SocialFacebook:  supplier.SocialFacebook,
			SocialTwitter:   supplier.SocialTwitter,
			SocialInstagram: supplier.SocialInstagram,
			SocialYoutube:   supplier.SocialYoutube,
			SocialLinkedin:  supplier.SocialLinkedin,
			SocialWebsite:   supplier.SocialWebsite,
			ID:              id,
		}
		if err := store.queries.UpdateSupplierNoImage(ctx, params); err != nil {
			return fmt.Errorf("error while updating the supplier: %w", err)
		}
	}
	return nil
}

func (store *SupplierStoreImpl) DeleteSupplier(ctx context.Context, id string) error {
	if err := store.queries.DeleteSupplier(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting supplier: %w", err)
	}
	return nil
}
