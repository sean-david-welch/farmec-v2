package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type MachineStore interface {
	GetMachines(id string) ([]types.Machine, error)
	GetMachineById(id string) (*types.Machine, error)
	CreateMachine(machine *types.Machine) error
	UpdateMachine(id string, machine *types.Machine) error
	DeleteMachine(id string) error
}

type MachineStoreImpl struct {
	database *sql.DB
}

func NewMachineStore(database *sql.DB) *MachineStoreImpl {
	return &MachineStoreImpl{database: database}
}

func ScanMachine(row interface{}, machine *types.Machine) error {
	var scanner interface {
		Scan(dest ...interface{}) error
	}

	switch value := row.(type) {
	case *sql.Row:
		scanner = value
	case *sql.Rows:
		scanner = value
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return scanner.Scan(&machine.ID, &machine.SupplierID, &machine.Name, &machine.MachineImage, &machine.Description, &machine.MachineLink, &machine.Created)
}

func (store *MachineStoreImpl) GetMachines(id string) ([]types.Machine, error) {
	var machines []types.Machine

	query := `SELECT * FROM "Machine" WHERE "supplier_id" = ?`
	rows, err := store.database.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var machine types.Machine

		if err := ScanMachine(rows, &machine); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		machines = append(machines, machine)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return machines, nil
}

func (store *MachineStoreImpl) GetMachineById(id string) (*types.Machine, error) {
	query := `SELECT * FROM "Machine" WHERE id = ?`
	row := store.database.QueryRow(query, id)

	var machine types.Machine

	if err := ScanMachine(row, &machine); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &machine, nil
}

func (store *MachineStoreImpl) CreateMachine(machine *types.Machine) error {
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
