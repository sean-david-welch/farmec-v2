package stores

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type WarrantyStore interface {
	GetWarranties() ([]types.DealerOwnerInfo, error)
	GetWarrantyById(id string) (*types.WarrantyClaim, []types.PartsRequired, error)
	CreateWarranty(warranty *types.WarrantyClaim, parts []types.PartsRequired) error
	UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error
	DeleteWarranty(id string) error
}

type WarrantyStoreImpl struct {
	database *sql.DB
}

func NewWarrantyStore(database *sql.DB) *WarrantyStoreImpl {
	return &WarrantyStoreImpl{database: database}
}

func (store *WarrantyStoreImpl) GetWarranties() ([]types.DealerOwnerInfo, error) {
	var warranties []types.DealerOwnerInfo

	query := `SELECT "id", "dealer", "owner_name" FROM "WarrantyClaim" ORDER BY "created" DESC`
	rows, err := store.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var warranty types.DealerOwnerInfo

		if err := rows.Scan(&warranty.ID, &warranty.Dealer, &warranty.OwnerName); err != nil {
			return nil, fmt.Errorf("error occurred while interating over rows: %w", err)
		}

		warranties = append(warranties, warranty)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred after interating over rows: %w", err)
	}

	return warranties, nil
}

func (store *WarrantyStoreImpl) GetWarrantyById(id string) (*types.WarrantyClaim, []types.PartsRequired, error) {
	var warranty types.WarrantyClaim
	var parts []types.PartsRequired

	warrantyQuery := `SELECT * FROM "WarrantyClaim" WHERE "id" = ?`

	row := store.database.QueryRow(warrantyQuery, id)

	if err := row.Scan(&warranty.ID, &warranty.Dealer, &warranty.DealerContact, &warranty.OwnerName,
		&warranty.MachineModel, &warranty.SerialNumber, &warranty.InstallDate, &warranty.FailureDate, &warranty.RepairDate,
		&warranty.FailureDetails, &warranty.RepairDetails, &warranty.LabourHours, &warranty.CompletedBy, &warranty.Created,
	); err != nil {
		return nil, nil, fmt.Errorf("error while querying database: %w", err)
	}

	partsQuery := `SELECT * FROM "PartsRequired" WHERE "warranty_id" = ?`
	rows, err := store.database.Query(partsQuery, id)
	if err != nil {
		return nil, nil, fmt.Errorf("error querying parts required from database: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var part types.PartsRequired

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

func (store *WarrantyStoreImpl) CreateWarranty(warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	warranty.ID = uuid.NewString()
	warranty.Created = time.Now().String()

	transaction, err := store.database.Begin()
	if err != nil {
		return err
	}

	warrantyQuery := `INSERT INTO "WarrantyClaim"
	(id, dealer, dealer_contact, owner_name, machine_model, serial_number, install_date, failure_date, repair_date, failure_details, repair_details, labour_hours, completed_by, created)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = transaction.Exec(
		warrantyQuery, warranty.ID, warranty.Dealer, warranty.DealerContact, warranty.OwnerName,
		warranty.MachineModel, warranty.SerialNumber, warranty.InstallDate, warranty.FailureDate, warranty.RepairDate,
		warranty.FailureDetails, warranty.RepairDetails, warranty.LabourHours, warranty.CompletedBy, warranty.Created,
	)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	partsQuery := `INSERT INTO "PartsRequired"
	(id, "warranty_id", part_number, quantity_needed, invoice_number, description) VALUES (?, ?, ?, ?, ?, ?)`

	for _, part := range parts {
		part.ID = uuid.NewString()

		_, err := transaction.Exec(partsQuery, part.ID, warranty.ID, part.PartNumber, part.QuantityNeeded, part.InvoiceNumber, part.Description)
		if err != nil {
			err := transaction.Rollback()
			if err != nil {
				return err
			}
			return err
		}

	}

	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (store *WarrantyStoreImpl) UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	transaction, err := store.database.Begin()
	if err != nil {
		return err
	}

	warrantyQuery := `UPDATE "WarrantyClaim" SET
    dealer = ?, dealer_contact = ?, owner_name = ?, machine_model = ?, serial_number = ?,
    install_date = ?, failure_date = ?, repair_date = ?, failure_details = ?, repair_details = ?,
    labour_hours = ?, completed_by = ?, created = ? WHERE id = ?`

	_, err = transaction.Exec(
		warrantyQuery, id, warranty.Dealer, warranty.DealerContact, warranty.OwnerName,
		warranty.MachineModel, warranty.SerialNumber, warranty.InstallDate, warranty.FailureDate, warranty.RepairDate,
		warranty.FailureDetails, warranty.RepairDetails, warranty.LabourHours, warranty.CompletedBy, warranty.Created,
	)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	deleteQuery := `DELETE FROM "PartsRequired" WHERE warranty_id = ?`
	_, err = transaction.Exec(deleteQuery, id)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	partsQuery := `INSERT INTO "PartsRequired"
	(id, warranty_id, part_number, quantity_needed, invoice_number, description) VALUES (?, ?, ?, ?, ?, ?)`

	for _, part := range parts {
		part.ID = uuid.NewString()

		_, err := transaction.Exec(partsQuery, part.ID, warranty.ID, part.PartNumber, part.QuantityNeeded, part.InvoiceNumber, part.Description)
		if err != nil {
			err := transaction.Rollback()
			if err != nil {
				return err
			}
			return err
		}
	}

	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (store *WarrantyStoreImpl) DeleteWarranty(id string) error {
	transaction, err := store.database.Begin()
	if err != nil {
		return err
	}

	deleteParts := `DELETE FROM "PartsRequired" WHERE "warranty_id" = ?`
	_, err = transaction.Exec(deleteParts, id)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		log.Printf("error occurred while deleting parts from warranty: %v", err)
		return err
	}

	deleteWarranty := `DELETE FROM "WarrantyClaim" WHERE "id" = ?`
	_, err = transaction.Exec(deleteWarranty, id)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		log.Printf("error occurred while deleting warranty claim: %v", err)
		return err
	}

	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil
}
