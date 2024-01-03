package services

import (
	"database/sql"

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
	rows, error := service.db.Query(query)

	if error != nil {return nil, error}

	defer rows.Close()


	for rows.Next() {
		var supplier models.Supplier

		if error := rows.Scan(&supplier.ID, &supplier.Name); error != nil {
			return nil, error
		}
		suppliers = append(suppliers, supplier)
	}

	if error = rows.Err(); error != nil {
		return nil, error
	}

	return suppliers, nil
}