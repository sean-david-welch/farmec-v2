package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

// type Sparepart struct {
//     ID              string `json:"id"`
//     SupplierID      string `json:"supplierId"`
//     Name            string `json:"name"`
//     PartsImage      string `json:"parts_image"`
//     SparePartsLink  string `json:"spare_parts_link"`
//     PdfLink         string `json:"pdf_link"`
// }

type PartsRepository struct {
	database *sql.DB
}

func NewPartsRepository(database *sql.DB) *PartsRepository {
	return &PartsRepository{database: database}
}

func ScanParts(row interface{}, part *types.Sparepart) error {
	var scanner interface {
		Scan(dest ...interface{}) error
	}

	switch value := row.(type) {
	case *sql.Row:
		scanner = value
	case *sql.Rows:
		scanner = value
	default: return fmt.Errorf("unsupported type: %T", value)
	}

	return scanner.Scan(&part.ID, &part.SupplierID, &part.Name, &part.PartsImage, &part.SparePartsLink, &part.PdfLink)
}

func (repository *PartsRepository) GetParts(id string) ([]types.Sparepart, error) {
	var parts []types.Sparepart

	query := `SELECT * FROM "SpareParts" WHERE "supplierId" = $1`
	rows, err := repository.database.Query(query, id); if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var part types.Sparepart

		if err := ScanParts(rows, &part); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		parts = append(parts, part)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return parts, nil
}

func (repository *PartsRepository) GetPartById(id string) (*types.Sparepart, error) {
	query := `SELECT * FROM "SpareParts" WHERE id = $1`
	row := repository.database.QueryRow(query, id)

	var part types.Sparepart

	if err := ScanParts(row, &part); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &part, nil
}

func (repository *PartsRepository) CreatePart(part *types.Sparepart) error {
	part.ID = uuid.NewString()
	
	query := `INSERT INTO "SpareParts" (ID, supplierID, name, partsImage, sparePartsLink, pdfLink)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repository.database.Exec(query, part.ID, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink, part.PdfLink)

	if err != nil {
		return fmt.Errorf("error creating spare part: %w", err)
	}

	return nil
}

func (repository *PartsRepository) UpdatePart(id string, part *types.Sparepart) error {
	query := `UPDATE "SpareParts"
	SET supplierID = $1, name = $2, partsImage = $3, sparePartsLink  = $4, pdfLink = $5
	WHERE ID = $6`

	_, err := repository.database.Exec(query, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink, part.PdfLink)

	if err != nil {
		return fmt.Errorf("error updating part: %w", err)
	}

	return nil
}

func (repository *PartsRepository) DeletePart(id string) error {
	query := `DELETE FROM "SpareParts" WHERE id = $1`

	_, err := repository.database.Exec(query, id); if err != nil {
		return fmt.Errorf("error deleting part: %w", err)
	}
	
	return nil
}