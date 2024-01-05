package services

import (
	"database/sql"
	"fmt"

	"github.com/sean-david-welch/farmec-v2/server/models"
)

type SupplierService struct {
	db *sql.DB
}

func NewSupplierService(db *sql.DB) *SupplierService {
	return &SupplierService{db: db}
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

    return scanner.Scan(&supplier.ID, &supplier.Name, &supplier.Description, &supplier.LogoImage, &supplier.MarketingImage, &supplier.SocialFacebook, &supplier.SocialInstagram, &supplier.SocialLinkedin, &supplier.SocialTwitter, &supplier.SocialYoutube, &supplier.SocialWebsite, &supplier.Created)
}

// pointer receiver avoids copying entire struct
func (service *SupplierService) GetSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier

	query := `SELECT * FROM "Supplier"`
	rows, err := service.db.Query(query)

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

func (service *SupplierService) GetSupplierByID(id string) (*models.Supplier, error) {
	query := `SELECT * FROM "Supplier" WHERE id = $1`
	row := service.db.QueryRow(query, id)

	var supplier models.Supplier

	if err := scanSupplier(row, &supplier); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning rown: %w", err)
	}

	return &supplier, nil
}

func (service *SupplierService) CreateSupplier(supplier *models.Supplier) error {
	query := `INSERT INTO "Supplier" 
				(name, description, logo_image, marketing_image, social_facebook, social_instagram, social_linkedin, social_twitter, social_youtube, social_website, created) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := service.db.Exec(query, supplier.Name, supplier.Description, supplier.LogoImage, supplier.MarketingImage, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, supplier.Created)

	if err != nil {
		return fmt.Errorf("error creating supplier: %w", err)
	}

	return nil
}

func (service *SupplierService) UpdateSupplier(id string, supplier *models.Supplier) error {
    query := `UPDATE "Supplier" SET 
                name = $1, 
                description = $2, 
                logo_image = $3, 
                marketing_image = $4, 
                social_facebook = $5, 
                social_instagram = $6, 
                social_linkedin = $7, 
                social_twitter = $8, 
                social_youtube = $9, 
                social_website = $10, 
                created = $11 
                WHERE id = $12`

    _, err := service.db.Exec(query, supplier.Name, supplier.Description, supplier.LogoImage, supplier.MarketingImage, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, supplier.Created, id)

    if err != nil {
        return fmt.Errorf("error updating supplier: %w", err)
    }

    return nil
}

func (service *SupplierService) DeleteSupplier(id string) error {
	query := `DELETE FROM "Supplier" WHERE id = $1`

	_, err := service.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deletting supplier: %w", err)
	}

	return nil 
}