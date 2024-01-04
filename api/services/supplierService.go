package services

import (
	"database/sql"
	"fmt"

	"githib.com/sean-david-welch/Farmec-Astro/api/models"
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