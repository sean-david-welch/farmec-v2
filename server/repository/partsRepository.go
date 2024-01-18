package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PartsRepository interface {
	GetParts(id string) ([]types.Sparepart, error) 
	GetPartById(id string) (*types.Sparepart, error) 
	CreatePart(part *types.Sparepart) error 
	UpdatePart(id string, part *types.Sparepart) error 
	DeletePart(id string) error 
}

type PartsRepositoryImpl struct {
	database *sql.DB
}

func NewPartsRepository(database *sql.DB) *PartsRepositoryImpl {
	return &PartsRepositoryImpl{database: database}
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

func (repository *PartsRepositoryImpl) GetParts(id string) ([]types.Sparepart, error) {
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

func (repository *PartsRepositoryImpl) GetPartById(id string) (*types.Sparepart, error) {
	query := `SELECT * FROM "SpareParts" WHERE id = $1`
	row := repository.database.QueryRow(query, id)

	var part types.Sparepart

	if err := ScanParts(row, &part); err != nil {

				if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &part, nil
}

func (repository *PartsRepositoryImpl) CreatePart(part *types.Sparepart) error {
	part.ID = uuid.NewString()
	
	query := `INSERT INTO "SpareParts" (ID, supplierID, name, partsImage, sparePartsLink, pdfLink)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repository.database.Exec(query, part.ID, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink, part.PdfLink)

	if err != nil {
		return fmt.Errorf("error creating spare part: %w", err)
	}

	return nil
}

func (repository *PartsRepositoryImpl) UpdatePart(id string, part *types.Sparepart) error {
	query := `UPDATE "SpareParts"
	SET supplierID = $1, name = $2, partsImage = $3, sparePartsLink  = $4, pdfLink = $5
	WHERE ID = $6`

	_, err := repository.database.Exec(query, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink, part.PdfLink)

	if err != nil {
		return fmt.Errorf("error updating part: %w", err)
	}

	return nil
}

func (repository *PartsRepositoryImpl) DeletePart(id string) error {
	query := `DELETE FROM "SpareParts" WHERE id = $1`

	_, err := repository.database.Exec(query, id); if err != nil {
		return fmt.Errorf("error deleting part: %w", err)
	}
	
	return nil
}