package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type WarrantyRepo interface {
	GetWarranties(ctx context.Context) ([]types.DealerOwnerInfo, error)
	GetWarrantyById(ctx context.Context, id string) (*db.WarrantyClaim, []db.PartsRequired, error)
	CreateWarranty(ctx context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	UpdateWarranty(ctx context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error
	DeleteWarranty(ctx context.Context, id string) error
}

type WarrantyRepoImpl struct {
	queries *db.Queries
	db      *sql.DB
}

func NewWarrantyRepo(sqlDB *sql.DB) *WarrantyRepoImpl {
	queries := db.New(sqlDB)
	return &WarrantyRepoImpl{queries: queries, db: sqlDB}
}

func (store *WarrantyRepoImpl) GetWarranties(ctx context.Context) ([]types.DealerOwnerInfo, error) {
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

func (store *WarrantyRepoImpl) GetWarrantyById(ctx context.Context, id string) (*db.WarrantyClaim, []db.PartsRequired, error) {
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

func (store *WarrantyRepoImpl) CreateWarranty(ctx context.Context, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)

	qtx := store.queries.WithTx(tx)

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
	if err := qtx.CreateWarranty(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating warranty: %w", err)
	}

	for _, part := range parts {
		part.ID = uuid.NewString()
		params := db.CreatePartsRequiredParams{
			ID:             part.ID,
			WarrantyID:     warranty.ID,
			PartNumber:     part.PartNumber,
			QuantityNeeded: part.QuantityNeeded,
			InvoiceNumber:  part.InvoiceNumber,
			Description:    part.Description,
		}
		if err := qtx.CreatePartsRequired(ctx, params); err != nil {
			return fmt.Errorf("an error occurred while creating the part: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
func (store *WarrantyRepoImpl) UpdateWarranty(ctx context.Context, id string, warranty *db.WarrantyClaim, parts []db.PartsRequired) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)
	qtx := store.queries.WithTx(tx)

	params := db.UpdateWarrantyParams{
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
		ID:             id,
	}
	if err := qtx.UpdateWarranty(ctx, params); err != nil {
		return fmt.Errorf("error occurred while updating warranty: %w", err)
	}
	if err = qtx.DeletePartsRequired(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting parts required: %w", err)
	}

	for _, part := range parts {
		part.ID = uuid.NewString()
		params := db.CreatePartsRequiredParams{
			ID:             part.ID,
			WarrantyID:     warranty.ID,
			PartNumber:     part.PartNumber,
			QuantityNeeded: part.QuantityNeeded,
			InvoiceNumber:  part.InvoiceNumber,
			Description:    part.Description,
		}
		if err := qtx.CreatePartsRequired(ctx, params); err != nil {
			return fmt.Errorf("an error occurred while creating the part: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (store *WarrantyRepoImpl) DeleteWarranty(ctx context.Context, id string) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)
	qtx := store.queries.WithTx(tx)

	if err = qtx.DeletePartsRequired(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting parts required: %w", err)
	}
	if err = qtx.DeleteWarranty(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting warranty: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
