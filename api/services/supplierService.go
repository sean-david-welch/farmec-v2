package services

import (
	"database/sql"
	"log"

	"githib.com/sean-david-welch/Farmec-Astro/api/models"
)

type SupplierService struct {
	db *sql.DB
}

func NewSupplierService(db *sql.DB) *SupplierService {
	return &SupplierService{db: db}
}

func (service SupplierService) GetSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier

	query := `SELECT * FROM "Supplier"`
	rows, err := service.db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier models.Supplier
		if err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Description, &supplier.LogoImage, &supplier.MarketingImage, &supplier.SocialFacebook, &supplier.SocialInstagram, &supplier.SocialLinkedin, &supplier.SocialTwitter, &supplier.SocialYoutube, &supplier.SocialWebsite, &supplier.Created); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating over rows: %v\n", err)
		return nil, err
	}

	return suppliers, nil
}