package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
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
	db *sql.DB
}

func NewPartsRepository(db *sql.DB) *PartsRepository {
	return &PartsRepository{db: db}
}

func ScanParts(row interface{}, part *models.Sparepart) error {
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

func (repository *PartsRepository) GetParts(id string) ([]models.Sparepart, error) {
	var parts []models.Sparepart

	query := `SELECT * FROM "SpareParts" WHERE "supplierId" = $1`
	rows, err := repository.db.Query(query); if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var part models.Sparepart

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

func (repository *PartsRepository) GetPartById(id string) (*models.Sparepart, error) {
	query := `SELECT * FROM "SpareParts" WHERE id = $1`
	row := repository.db.QueryRow(query, id)

	var part models.Sparepart

	if err := ScanParts(row, &part); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &part, nil
}

func (repository *PartsRepository) CreatePart(part *models.Sparepart) error {
	part.ID = uuid.NewString()
	
	query := `INSERT INTO "SpareParts" (ID, SupplierID, Name, PartsImage, SparePartsLink, PdfLink)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repository.db.Exec(query, part.ID, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink, part.PdfLink)

	if err != nil {
		return fmt.Errorf("error creating spare part: %w", err)
	}

	return nil
}

func (repository *PartsRepository) UpdatePart(id string, part *models.Sparepart) error {
	query := `UPDATE "SpareParts"
	SET SupplierID = $1, Name = $2, PartsImage = $3, SparePartsLink  = $4, PdfLink = $5
	WHERE ID = $6`

	_, err := repository.db.Exec(query, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink, part.PdfLink)

	if err != nil {
		return fmt.Errorf("error updating part: %w", err)
	}

	return nil
}

func (repository *PartsRepository) DeletePart(id string) error {
	query := `DELETE FROM "SpareParts" WHERE id = $1`

	_, err := repository.db.Exec(query, id); if err != nil {
		return fmt.Errorf("error deleting part: %w", err)
	}
	
	return nil
}