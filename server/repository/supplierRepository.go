package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func scanSupplier(row interface{}, supplier *models.Supplier) error {
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
func (repository *SupplierRepository) GetSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier

	query := `SELECT * FROM "Supplier"`
	rows, err := repository.db.Query(query)

    if err != nil {
        return nil, fmt.Errorf("error executing query: %w", err)
    }
	defer rows.Close()

	for rows.Next() {
		var supplier models.Supplier

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

func (repository *SupplierRepository) CreateSupplier(supplier *models.Supplier) error {

	supplier.ID = uuid.NewString()
	supplier.Created = time.Now()

	log.Printf("Creating supplier: %+v", supplier)


	query := `INSERT INTO "Supplier" 
	(id, name, logo_image, marketing_image, description, social_facebook, social_instagram, social_linkedin, social_twitter, social_youtube, social_website, created) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := repository.db.Exec(query, supplier.ID, supplier.Name, supplier.LogoImage, supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, supplier.Created)

	if err != nil {
		return fmt.Errorf("error creating supplier: %w", err)
	}

	return nil
}

func (repository *SupplierRepository) GetSupplierById(id string) (*models.Supplier, error) {
	query := `SELECT * FROM "Supplier" WHERE id = $1`
	row := repository.db.QueryRow(query, id)

	var supplier models.Supplier

	if err := scanSupplier(row, &supplier); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning rown: %w", err)
	}

	return &supplier, nil
}

func (repository *SupplierRepository) UpdateSupplier(id string, supplier *models.Supplier) error {
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

    _, err := repository.db.Exec(query, supplier.Name, supplier.LogoImage, supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, id)

    if err != nil {
        return fmt.Errorf("error updating supplier: %w", err)
    }

    return nil
}

func (repository *SupplierRepository) DeleteSupplier(id string) error {
	query := `DELETE FROM "Supplier" WHERE id = $1`

	_, err := repository.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deleting supplier: %w", err)
	}

	return nil 
}