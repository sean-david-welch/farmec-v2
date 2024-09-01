package stores

import (
	"context"
	"database/sql"
	"errors"
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

func (store *SupplierStoreImpl) GetSupplierById(id string) (*db.Supplier, error) {
	query := `SELECT * FROM "Supplier" WHERE id = ?`
	row := store.database.QueryRow(query, id)

	var supplier db.Supplier

	if err := scanSupplier(row, &supplier); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
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

func (store *SupplierStoreImpl) UpdateSupplier(id string, supplier *db.Supplier) error {
	query := `UPDATE "Supplier" SET 
                name = ?,  
                description = ?, 
                social_facebook = ?, 
                social_instagram = ?, 
                social_linkedin = ?, 
                social_twitter = ?, 
                social_youtube = ?, 
                social_website = ? 
                WHERE id = ?`
	args := []interface{}{supplier.Name, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, id}

	if supplier.LogoImage != "" && supplier.LogoImage != "null" && supplier.MarketingImage != "" && supplier.MarketingImage != "null" {
		query = `UPDATE "Supplier" SET 
		name = ?, 
		logo_image = ?, 
		marketing_image = ?, 
		description = ?, 
		social_facebook = ?, 
		social_instagram = ?, 
		social_linkedin = ?, 
		social_twitter = ?, 
		social_youtube = ?, 
		social_website = ? 
		WHERE id = ?`
		args = []interface{}{supplier.Name, supplier.LogoImage, supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, id}
	}

	_, err := store.database.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("error updating supplier: %w", err)
	}

	return nil
}

func (store *SupplierStoreImpl) DeleteSupplier(id string) error {
	query := `DELETE FROM "Supplier" WHERE id = ?`

	_, err := store.database.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deleting supplier: %w", err)
	}

	return nil
}
