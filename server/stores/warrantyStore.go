package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type WarrantyStore interface {
	GetWarranties(ctx context.Context) ([]types.DealerOwnerInfo, error)
	GetWarrantyById(ctx context.Context, id string) (*db.WarrantyClaim, []db.PartsRequired, error)
	CreateWarranty(ctx context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	UpdateWarranty(ctx context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	DeleteWarranty(ctx context.Context, id string) error
}

type WarrantyStoreImpl struct {
	queries *db.Queries
}

func NewWarrantyStore(sql *sql.DB) *WarrantyStoreImpl {
	queries := db.New(sql)
	return &WarrantyStoreImpl{queries: queries}
}

func (store *WarrantyStoreImpl) GetWarranties(ctx context.Context) ([]types.DealerOwnerInfo, error) {
	warranties, err := store.queries.GetWarranties(ctx)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting warranties: %w", err)
	}

	var result []types.DealerOwnerInfo
	for _, warranty := range warranties {
		result = append(result, types.DealerOwnerInfo{
			ID:        warranty.ID,
			Dealer:    warranty.Dealer,
			OwnerName: warranty.OwnerName,
		})
	}
	return result, nil
}

func (store *WarrantyStoreImpl) GetWarrantyById(ctx context.Context, id string) (*db.WarrantyClaim, []db.PartsRequired, error) {
	rows, err := store.queries.GetWarrantyByID(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("error occurred while getting warranty from the db: %w", err)
	}

	var warranty *db.WarrantyClaim
	var parts []db.PartsRequired

	for _, row := range rows {
		if warranty == nil {
			warranty = &db.WarrantyClaim{
				ID:             row.ID,
				Dealer:         row.Dealer,
				DealerContact:  row.DealerContact,
				OwnerName:      row.OwnerName,
				MachineModel:   row.MachineModel,
				SerialNumber:   row.SerialNumber,
				InstallDate:    row.InstallDate,
				FailureDate:    row.FailureDate,
				RepairDate:     row.RepairDate,
				FailureDetails: row.FailureDetails,
				RepairDetails:  row.RepairDetails,
				LabourHours:    row.LabourHours,
				CompletedBy:    row.CompletedBy,
				Created:        row.Created,
			}
		}
		if row.PartID.Valid {
			parts = append(parts, db.PartsRequired{
				ID:             row.PartID.String,
				WarrantyID:     row.ID,
				PartNumber:     row.PartNumber,
				QuantityNeeded: row.QuantityNeeded.String,
				InvoiceNumber:  row.InvoiceNumber,
				Description:    row.Description,
			})
		}
	}

	if warranty == nil {
		return nil, nil, fmt.Errorf("warranty with id %s not found", id)
	}
	return warranty, parts, nil
}

func (store *WarrantyStoreImpl) CreateWarranty(ctx context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	warranty.ID = uuid.NewString()
	warranty.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}

	params := db.CreateWarrantyParams{
		ID:             warranty.ID,
		Dealer:         warranty.Dealer,
		DealerContact:  warranty.DealerContact,
		OwnerName:      warranty.OwnerName,
		MachineModel:   warranty.MachineModel,
		SerialNumber:   warranty.SerialNumber,
		InstallDate:    warranty.InstallDate,
		FailureDate:    warranty.FailureDate,
		RepairDate:     warranty.RepairDate,
		FailureDetails: warranty.FailureDetails,
		RepairDetails:  warranty.RepairDetails,
		LabourHours:    warranty.LabourHours,
		CompletedBy:    warranty.CompletedBy,
		Created:        warranty.Created,
	}
	if err := store.queries.CreateWarranty(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating warranty: %w", err)
	}

	for _, part := range parts {
		part.ID = uuid.NewString()
		params := db.CreatePartsRequiredParams{
			ID:             part.ID,
			WarrantyID:     part.WarrantyID,
			PartNumber:     part.PartNumber,
			QuantityNeeded: part.QuantityNeeded,
			InvoiceNumber:  part.InvoiceNumber,
			Description:    part.Description,
		}
		if err := store.queries.CreatePartsRequired(ctx, params); err != nil {
			return fmt.Errorf("an error occurred while creating the part: %w", err)
		}
	}
	return nil
}

func (store *WarrantyStoreImpl) UpdateWarranty(id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
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
