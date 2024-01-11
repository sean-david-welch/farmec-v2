package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

// type Machine struct {
//     ID            string `json:"id"`
//     SupplierID    string `json:"supplierId"`
//     Name          string `json:"name"`
//     MachineImage  string `json:"machine_image"`
//     Description   *string `json:"description"`
//     MachineLink   *string `json:"machine_link"`
// 	   Created       time.Time `json:"created"`
// }

type MachineRepository struct {
	db *sql.DB
}

func NewMachineRepository(db *sql.DB) *MachineRepository {
	return &MachineRepository{db: db}
}

func ScanMachine(row interface{}, machine *models.Machine) error {
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

	return scanner.Scan(&machine.ID, &machine.SupplierID, &machine.Name, &machine.MachineImage, &machine.Description, &machine.MachineLink, &machine.Created)
}

func (repository *MachineRepository) GetMachines(id string) ([]models.Machine, error) {
	var machines []models.Machine

	query := `SELECT * FROM "Machine" WHERE "supplierId" = $1`
	rows, err := repository.db.Query(query, id); if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	for rows.Next(){
		var machine models.Machine

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

func (repository *MachineRepository) GetMachineById(id string) (*models.Machine, error) {
	query := `SELECT * FROM "Machine" WHERE id = $1`
	row := repository.db.QueryRow(query, id)

	var machine models.Machine

	if err := ScanMachine(row, &machine); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &machine, nil
}

func (repository *MachineRepository) CreateMachine(machine *models.Machine) error {
	machine.ID = uuid.NewString()
	machine.Created = time.Now()

	query := `INSERT INTO Machine (id, supplierId, name, machineImage, description, machineLink, created) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`	

	_, err := repository.db.Exec(query, machine.ID, machine.SupplierID, machine.Name, machine.MachineImage, machine.Description, machine.MachineLink, machine.Created)

	if err != nil {
		return fmt.Errorf("error creating machine: %w", err)
	}

	return nil
}

func (repository *MachineRepository) UpdateMachine(id string, machine *models.Machine) error {
	query := `UPDATE Machine 
	SET supplierID = $1, name = $2, machineImage = $3, description = $4, machineLink = $5
	WHERE ID = $6`

	_, err := repository.db.Exec(query, machine.SupplierID, machine.Name, machine.MachineImage, machine.Description, machine.MachineLink, id)
	
	if err != nil {
		return fmt.Errorf("error updating machine: %w", err)
	}

	return nil
}

func (repository *MachineRepository) DeleteMachine(id string) error {
	query := `DELETE FROM "Machine" WHERE id = $1`

	_, err := repository.db.Exec(query, id); if err != nil {
		return fmt.Errorf("error deleting machine: %w", err)
	}

	return nil
}
