package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
)

type SupplierRepo interface {
	GetSuppliers(ctx context.Context) ([]db.Supplier, error)
	GetSupplierVidoes(ctx context.Context, id string) ([]db.Video, error)
	GetSupplierMachines(ctx context.Context, id string) ([]db.Machine, error)
	GetSupplierById(ctx context.Context, id string) (*db.Supplier, error)
	CreateSupplier(ctx context.Context, supplier *db.Supplier) error
	UpdateSupplier(ctx context.Context, id string, supplier *db.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error
}

type SupplierRepoImpl struct {
	queries *db.Queries
}

func NewSupplierRepo(sql *sql.DB) *SupplierRepoImpl {
	queries := db.New(sql)
	return &SupplierRepoImpl{queries: queries}
}

func (repo *SupplierRepoImpl) GetSuppliers(ctx context.Context) ([]db.Supplier, error) {
	suppliers, err := repo.queries.GetSuppliers(ctx)
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

func (repo *SupplierRepoImpl) GetSupplierVidoes(ctx context.Context, id string) ([]db.Video, error) {
	videos, err := repo.queries.GetVideosBySupplierID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting video from the db: %w", err)
	}

	var result []db.Video
	for _, video := range videos {
		result = append(result, db.Video{
			ID:           video.ID,
			SupplierID:   video.SupplierID,
			WebUrl:       video.WebUrl,
			Title:        video.Title,
			Description:  video.Description,
			VideoID:      video.VideoID,
			ThumbnailUrl: video.ThumbnailUrl,
			Created:      video.Created,
		})
	}
	return result, nil
}

func (repo *SupplierRepoImpl) GetSupplierMachines(ctx context.Context, id string) ([]db.Machine, error) {
	machines, err := repo.queries.GetMachinesBySupplierID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting machines from the db: %w", err)
	}

	var result []db.Machine
	for _, machine := range machines {
		result = append(result, db.Machine{
			ID:           machine.ID,
			SupplierID:   machine.SupplierID,
			Name:         machine.Name,
			MachineImage: machine.MachineImage,
			Description:  machine.Description,
			MachineLink:  machine.MachineLink,
			Created:      machine.Created,
		})
	}
	return result, nil
}

func (repo *SupplierRepoImpl) GetSupplierById(ctx context.Context, id string) (*db.Supplier, error) {
	supplierRow, err := repo.queries.GetSupplierByID(ctx, id)
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

func (repo *SupplierRepoImpl) CreateSupplier(ctx context.Context, supplier *db.Supplier) error {
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
	if err := repo.queries.CreateSupplier(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating supplier: %w", err)
	}
	return nil
}

func (repo *SupplierRepoImpl) UpdateSupplier(ctx context.Context, id string, supplier *db.Supplier) error {
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
		if err := repo.queries.UpdateSupplier(ctx, params); err != nil {
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
		if err := repo.queries.UpdateSupplierNoImage(ctx, params); err != nil {
			return fmt.Errorf("error while updating the supplier: %w", err)
		}
	}
	return nil
}

func (repo *SupplierRepoImpl) DeleteSupplier(ctx context.Context, id string) error {
	if err := repo.queries.DeleteSupplier(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting supplier: %w", err)
	}
	return nil
}
