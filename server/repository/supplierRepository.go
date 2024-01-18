package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type SupplierRepository struct {
	database *sql.DB
}

func NewSupplierRepository(database *sql.DB) *SupplierRepository {
	return &SupplierRepository{database: database}
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

    return scanner.Scan(&supplier.ID, &supplier.Name,  &supplier.LogoImage, &supplier.MarketingImage, &supplier.Description, &supplier.SocialFacebook, &supplier.SocialInstagram, &supplier.SocialLinkedin, &supplier.SocialTwitter, &supplier.SocialYoutube, &supplier.SocialWebsite, &supplier.Created)
}

// pointer receiver avoids copying entire struct
func (repository *SupplierRepository) GetSuppliers() ([]types.Supplier, error) {
	var suppliers []types.Supplier

	query := `SELECT * FROM "Supplier"`
	rows, err := repository.database.Query(query)

    if err != nil {
        return nil, fmt.Errorf("error executing query: %w", err)
    }
	defer rows.Close()

	for rows.Next() {
		var supplier types.Supplier

		// scans rows and copys them to supplier struct
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

func (repository *SupplierRepository) CreateSupplier(supplier *types.Supplier) error {

	supplier.ID = uuid.NewString()
	supplier.Created = time.Now()

	query := `INSERT INTO "Supplier" 
	(id, name, logo_image, marketing_image, description, social_facebook, social_instagram, social_linkedin, social_twitter, social_youtube, social_website, created) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := repository.database.Exec(query, supplier.ID, supplier.Name, supplier.LogoImage, supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, supplier.Created)

	if err != nil {
		return fmt.Errorf("error creating supplier: %w", err)
	}

	return nil
}

func (repository *SupplierRepository) GetSupplierById(id string) (*types.Supplier, error) {
	query := `SELECT * FROM "Supplier" WHERE id = $1`
	row := repository.database.QueryRow(query, id)

	var supplier types.Supplier

	if err := scanSupplier(row, &supplier); err != nil {

				if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &supplier, nil
}

func (repository *SupplierRepository) UpdateSupplier(id string, supplier *types.Supplier) error {
    query := `UPDATE "Supplier" SET 
                name = $1, 
                logo_image = $2, 
                marketing_image = $3, 
                description = $4, 
                social_facebook = $5, 
                social_instagram = $6, 
                social_linkedin = $7, 
                social_twitter = $8, 
                social_youtube = $9, 
                social_website = $10 
                WHERE id = $11`

    _, err := repository.database.Exec(query, supplier.Name, supplier.LogoImage, supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, id)

    if err != nil {
        return fmt.Errorf("error updating supplier: %w", err)
    }

    return nil
}

func (repository *SupplierRepository) DeleteSupplier(id string) error {
	query := `DELETE FROM "Supplier" WHERE id = $1`

	_, err := repository.database.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deleting supplier: %w", err)
	}

	return nil 
}