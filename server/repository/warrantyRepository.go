package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

type WarrantyRepository struct {
	database *sql.DB
}

func NewWarrantyRepository(database *sql.DB) *WarrantyRepository {
	return &WarrantyRepository{database: database}
}

func(repository*WarrantyRepository) GetWarranties() ([]models.DealerOwnerInfo, error) {
	var warranties []models.DealerOwnerInfo

    query := `SELECT "dealer", "ownerName" FROM "WarrantyClaim"`
	rows, err := repository.database.Query(query); if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var warranty models.DealerOwnerInfo

		if err := rows.Scan()
		err != nil {
			return nil, fmt.Errorf("error occurred while interating over rows: %w", err)
		}

		warranties = append(warranties, warranty)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred after interating over rows: %w", err)
	}

	return warranties, nil
}

func(repository*WarrantyRepository) GetWarrantyById(id string) (*models.WarrantyClaim, []models.PartsRequired, error) {
	var warranty models.WarrantyClaim
	var parts []models.PartsRequired

	warrantyQuery := `SELECT * FROM "WarrantyClaim" WHERE id = $1`

	row := repository.database.QueryRow(warrantyQuery, id)
	
	if err := row.Scan(&warranty.ID, &warranty.Dealer, &warranty.DealerContact, &warranty.OwnerName, 
		&warranty.MachineModel, &warranty.SerialNumber, &warranty.InstallDate, &warranty.FailureDate, &warranty.RepairDate, 
		&warranty.FailureDetails, &warranty.RepairDetails, &warranty.LabourHours, &warranty.CompletedBy, &warranty.Created,
	); err != nil {
		return nil, nil, fmt.Errorf("error while querying database: %w", err)
	}

	partsQuery := `SELECT * FROM "PartsRequired" WHERE warrantyId = $1`
	rows, err := repository.database.Query(partsQuery, id); if err != nil {
        return nil, nil, fmt.Errorf("error querying parts required from database: %w", err)
    }
    defer rows.Close()

	for rows.Next() {
		var part models.PartsRequired

		if err := rows.Scan(&part.ID, &part.WarrantyID, &part.PartNumber, &part.QuantityNeeded, &part.InvoiceNumber, &part.Description); err != nil {
			return nil, nil, fmt.Errorf("error while iterating over rows")
		}

		parts = append(parts, part)
	}

	if err := rows.Err(); err != nil {
        return nil, nil, fmt.Errorf("error iterating over parts required rows: %w", err)
    }

	return &warranty, parts, nil
}

func(repository*WarrantyRepository) CreateWarranty(warranty *models.WarrantyClaim, parts []models.PartsRequired) error {
	warranty.ID = uuid.NewString()
	warranty.Created = time.Now()

	transaction, err := repository.database.Begin(); if err != nil {
		return err
	}

    warrantyQuery := `INSERT INTO "WarrantyClaim" 
	(id, dealer, dealerContact, ownerName, machineModel, serialNumber, installDate, failureDate, repairDate, failureDetails, repairDetails, labourHours, completedBy, created) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

    _, err = transaction.Exec(
		warrantyQuery, warranty.ID, warranty.Dealer, warranty.DealerContact, warranty.OwnerName, 
		warranty.MachineModel, warranty.SerialNumber, warranty.InstallDate, warranty.FailureDate, warranty.RepairDate, 
		warranty.FailureDetails, warranty.RepairDetails, warranty.LabourHours, warranty.CompletedBy, warranty.Created,
	); if err != nil {
		transaction.Rollback()
		return err
	}

    partsQuery := `INSERT INTO "PartsRequired" 
	(id, warrantyID, partNumber, quantityNeeded, invoiceNumber, description) VALUES ($1, $2, $3, $4, $5, $6)`

	for _, part := range parts {
		part.ID = uuid.NewString()

        _, err := transaction.Exec(partsQuery, part.ID, warranty.ID, part.PartNumber, part.QuantityNeeded, part.InvoiceNumber, part.Description)
		if err != nil {
			transaction.Rollback()
			return err
		}

	}

	if err := transaction.Commit(); err != nil {
		return err
	}
			
	return nil		
}

func(repository*WarrantyRepository) UpdateWarranty(id string, warranty *models.WarrantyClaim, parts []models.PartsRequired) error {
    transaction, err := repository.database.Begin(); if err != nil {
        return err
    }

    warrantyQuery := `UPDATE "WarrantyClaim" SET 
    dealer = $2, dealerContact = $3, ownerName = $4, machineModel = $5, serialNumber = $6, 
    installDate = $7, failureDate = $8, repairDate = $9, failureDetails = $10, repairDetails = $11, 
    labourHours = $12, completedBy = $13, created = $14 WHERE id = $1`

    _, err = transaction.Exec(
        warrantyQuery, id, warranty.Dealer, warranty.DealerContact, warranty.OwnerName, 
        warranty.MachineModel, warranty.SerialNumber, warranty.InstallDate, warranty.FailureDate, warranty.RepairDate, 
        warranty.FailureDetails, warranty.RepairDetails, warranty.LabourHours, warranty.CompletedBy, warranty.Created,
    ); if err != nil {
        transaction.Rollback()
        return err
    }

	deleteQuery := `DELETE FROM "PartsRequired" WHERE warrantyId = $1`
    _, err = transaction.Exec(deleteQuery, id); if err != nil {
        transaction.Rollback()
        return err
    }

    partsQuery := `INSERT INTO "PartsRequired" 
	(id, warrantyID, partNumber, quantityNeeded, invoiceNumber, description) VALUES ($1, $2, $3, $4, $5, $6)`

	for _, part := range parts {
		part.ID = uuid.NewString()

        _, err := transaction.Exec(partsQuery, part.ID, warranty.ID, part.PartNumber, part.QuantityNeeded, part.InvoiceNumber, part.Description)
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	if err := transaction.Commit(); err != nil {
		return err
	}
            
    return nil		
}

func(repository*WarrantyRepository) DeleteWarranty(id string) error {
	transaction, err := repository.database.Begin(); if err != nil {
		return err
	}

	deleteParts := `DELETE FROM "PartsRequired" WHERE "warrantyId" = $1`
	_, err = transaction.Exec(deleteParts, id); if err != nil {
		transaction.Rollback()
		return err
	}
	
	deleteWarranty := `DELETE FROM "WarrantyClaim" WHERE "id" = $1`
	_, err = transaction.Exec(deleteWarranty, id); if err != nil {
		transaction.Rollback()
		return err
	}


	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil 
}