package stores

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type SupplierStore interface {
	GetSuppliers() ([]types.Supplier, error)
	CreateSupplier(supplier *types.Supplier) error
	GetSupplierById(id string) (*types.Supplier, error)
	UpdateSupplier(id string, supplier *types.Supplier) error
	DeleteSupplier(id string) error
}

type SupplierStoreImpl struct {
	database *sql.DB
}

func NewSupplierStore(database *sql.DB) *SupplierStoreImpl {
	return &SupplierStoreImpl{database: database}
}

func scanSupplier(row interface{}, supplier *types.Supplier) error {
	var scanner interface {
		Scan(dest ...interface{}) error
	}

	switch value := row.(type) {
	case *sql.Row:
		scanner = value
	case *sql.Rows:
		scanner = value
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return scanner.Scan(
		&supplier.ID, &supplier.Name, &supplier.LogoImage,
		&supplier.MarketingImage, &supplier.Description,
		&supplier.SocialFacebook, &supplier.SocialInstagram, &supplier.SocialLinkedin,
		&supplier.SocialTwitter, &supplier.SocialYoutube, &supplier.SocialWebsite, &supplier.Created,
	)
}

func (store *SupplierStoreImpl) GetSuppliers() ([]types.Supplier, error) {
	var suppliers []types.Supplier

	query := `SELECT * FROM "Supplier" ORDER BY created DESC`
	rows, err := store.database.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var supplier types.Supplier

		if err := scanSupplier(rows, &supplier); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		suppliers = append(suppliers, supplier)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return suppliers, nil
}

func (store *SupplierStoreImpl) CreateSupplier(supplier *types.Supplier) error {

	supplier.ID = uuid.NewString()
	supplier.Created = time.Now().String()

	query := `INSERT INTO "Supplier" 
	(id, name, logo_image, marketing_image, description, social_facebook, social_instagram, social_linkedin, social_twitter, social_youtube, social_website, created) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, supplier.ID, supplier.Name, supplier.LogoImage, supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, supplier.Created)

	if err != nil {
		return fmt.Errorf("error creating supplier: %w", err)
	}

	return nil
}

func (store *SupplierStoreImpl) GetSupplierById(id string) (*types.Supplier, error) {
	query := `SELECT * FROM "Supplier" WHERE id = ?`
	row := store.database.QueryRow(query, id)

	var supplier types.Supplier

	if err := scanSupplier(row, &supplier); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &supplier, nil
}

func (store *SupplierStoreImpl) UpdateSupplier(id string, supplier *types.Supplier) error {
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
