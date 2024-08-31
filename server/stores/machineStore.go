package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
)

type MachineStore interface {
	GetMachines(ctx context.Context, id string) ([]db.Machine, error)
	GetMachineById(ctx context.Context, id string) (*db.Machine, error)
	CreateMachine(ctx context.Context, machine *db.Machine) error
	UpdateMachine(ctx context.Context, id string, machine *db.Machine) error
	DeleteMachine(ctx context.Context, id string) error
}

type MachineStoreImpl struct {
	queries *db.Queries
}

func NewMachineStore(sql *sql.DB) *MachineStoreImpl {
	queries := db.New(sql)
	return &MachineStoreImpl{queries: queries}
}

func (store *MachineStoreImpl) GetMachines(ctx context.Context, id string) ([]db.Machine, error) {
	machines, err := store.queries.GetMachines(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying the database for machines: %w", err)
	}

	var result []db.Machine
	for _, machine := range machines {
		result = append(result, db.Machine{
			ID:           machine.ID,
			SupplierID:   machine.SupplierID,
			Name:         machine.Name,
			MachineImage: machine.MachineImage,
			Description:  machine.Description,
			MachineLink:  machine.MachineLink,
			Created:      machine.Created,
		})
	}

	return result, nil
}

func (store *MachineStoreImpl) GetMachineById(ctx context.Context, id string) (*db.Machine, error) {
	machine, err := store.queries.GetMachineByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying the database for machines: %w", err)
	}

	return &machine, nil
}

func (store *MachineStoreImpl) CreateMachine(ctx context.Context, machine *db.Machine) error {
	machine.ID = uuid.NewString()
	machine.Created = sql.NullString{String: time.Now().String(), Valid: true}

	params := db.CreateMachineParams{
		ID:           machine.ID,
		SupplierID:   machine.SupplierID,
		Name:         machine.Name,
		MachineImage: machine.MachineImage,
		Description:  machine.Description,
		MachineLink:  machine.MachineLink,
		Created:      machine.Created,
	}

	if err := store.queries.CreateMachine(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating a machine: %w", err)
	}

	return nil
}

func (store *MachineStoreImpl) UpdateMachine(ctx context.Context, id string, machine *db.Machine) error {
	if machine.MachineImage.Valid {
		params := db.UpdateMachineParams{
			SupplierID:   machine.SupplierID,
			Name:         machine.Name,
			MachineImage: machine.MachineImage,
			Description:  machine.Description,
			MachineLink:  machine.MachineLink,
			ID:           machine.ID,
		}
		if err := store.queries.UpdateMachine(ctx, params); err != nil {
			return fmt.Errorf("error ocurred while updating a machine with image: %w", err)
		}
	} else {
		params := db.UpdateMachineNoImageParams{
			SupplierID:  machine.SupplierID,
			Name:        machine.Name,
			Description: machine.Description,
			MachineLink: machine.MachineLink,
			ID:          machine.ID,
		}
		if err := store.queries.UpdateMachineNoImage(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updating the machine without image: %w", err)
		}
	}
	return nil
}

func (store *MachineStoreImpl) DeleteMachine(ctx context.Context, id string) error {
	if err := store.queries.DeleteMachine(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting machine: %w", err)
	}
	return nil
}
