package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
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

func (store *MachineStoreImpl) GetMachines(id string) ([]db.Machine, error) {
	ctx := context.Background()
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

func (store *MachineStoreImpl) GetMachineById(id string) (*db.Machine, error) {
	ctx := context.Background()
	machine, err := store.queries.GetMachineByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying the database for machines: %w", err)
	}

	return &machine, nil
}

func (store *MachineStoreImpl) CreateMachine(machine *types.Machine) error {
	ctx := context.Background()
	machine.ID = uuid.NewString()
	machine.Created = time.Now().String()

	query := `INSERT INTO "Machine" (id, supplier_id, name, machine_image, description, machine_link, created)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, machine.ID, machine.SupplierID, machine.Name, machine.MachineImage, machine.Description, machine.MachineLink, machine.Created)

	if err != nil {
		return fmt.Errorf("error creating machine: %w", err)
	}

	return nil
}

func (store *MachineStoreImpl) UpdateMachine(id string, machine *types.Machine) error {
	query := `UPDATE "Machine" SET supplier_id = ?, name = ?, description = ?, machine_link = ? WHERE ID = ?`
	args := []interface{}{machine.SupplierID, machine.Name, machine.Description, machine.MachineLink, id}

	if machine.MachineImage != "" && machine.MachineImage != "null" {
		query = `UPDATE "Machine" SET supplier_id = ?, name = ?, machine_image = ?, description = ?, machine_link = ? WHERE ID = ?`
		args = []interface{}{machine.SupplierID, machine.Name, machine.MachineImage, machine.Description, machine.MachineLink, id}
	}

	_, err := store.database.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("error updating machine: %w", err)
	}

	return nil
}

func (store *MachineStoreImpl) DeleteMachine(id string) error {
	query := `DELETE FROM "Machine" WHERE id = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting machine: %w", err)
	}

	return nil
}
