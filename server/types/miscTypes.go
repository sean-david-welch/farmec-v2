package types

import "github.com/sean-david-welch/farmec-v2/server/db"

type Carousel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Created string `json:"created"`
}

type DealerOwnerInfo struct {
	ID        string `json:"id"`
	Dealer    string `json:"dealer"`
	OwnerName string `json:"owner_name"`
}

type WarrantyClaim struct {
	ID             string  `json:"id"`
	Dealer         string  `json:"dealer"`
	DealerContact  *string `json:"dealer_contact"`
	OwnerName      *string `json:"owner_name"`
	MachineModel   *string `json:"machine_model"`
	SerialNumber   *string `json:"serial_number"`
	InstallDate    *string `json:"install_date"`
	FailureDate    *string `json:"failure_date"`
	RepairDate     *string `json:"repair_date"`
	FailureDetails *string `json:"failure_details"`
	RepairDetails  *string `json:"repair_details"`
	LabourHours    *string `json:"labour_hours"`
	CompletedBy    *string `json:"completed_by"`
	Created        string  `json:"created"`
}

type PartsRequired struct {
	ID             string `json:"id"`
	WarrantyID     string `json:"warranty_id"`
	PartNumber     string `json:"part_number"`
	QuantityNeeded string `json:"quantity_needed"`
	InvoiceNumber  string `json:"invoice_number"`
	Description    string `json:"description"`
}

type WarranrtyParts struct {
	Warranty *db.WarrantyClaim
	Parts    []db.PartsRequired
}

type WarrantyClaimPDF struct {
	Dealer         string  `json:"dealer"`
	DealerContact  *string `json:"dealer_contact"`
	OwnerName      *string `json:"owner_name"`
	MachineModel   *string `json:"machine_model"`
	SerialNumber   *string `json:"serial_number"`
	InstallDate    *string `json:"install_date"`
	FailureDate    *string `json:"failure_date"`
	RepairDate     *string `json:"repair_date"`
	FailureDetails *string `json:"failure_details"`
	RepairDetails  *string `json:"repair_details"`
	LabourHours    *string `json:"labour_hours"`
	CompletedBy    *string `json:"completed_by"`
}

type PartsRequiredPDF struct {
	PartNumber     string `json:"part_number"`
	QuantityNeeded string `json:"quantity_needed"`
	InvoiceNumber  string `json:"invoice_number"`
	Description    string `json:"description"`
}

type MachineRegistration struct {
	ID               string `json:"id"`
	DealerName       string `json:"dealer_name"`
	DealerAddress    string `json:"dealer_address"`
	OwnerName        string `json:"owner_name"`
	OwnerAddress     string `json:"owner_address"`
	MachineModel     string `json:"machine_model"`
	SerialNumber     string `json:"serial_number"`
	InstallDate      string `json:"install_date"`
	InvoiceNumber    string `json:"invoice_number"`
	CompleteSupply   bool   `json:"complete_supply"`
	PdiComplete      bool   `json:"pdi_complete"`
	PtoCorrect       bool   `json:"pto_correct"`
	MachineTestRun   bool   `json:"machine_test_run"`
	SafetyInduction  bool   `json:"safety_induction"`
	OperatorHandbook bool   `json:"operator_handbook"`
	Date             string `json:"date"`
	CompletedBy      string `json:"completed_by"`
	Created          string `json:"created"`
}

type MachineRegistrationPDF struct {
	DealerName       string `json:"dealer_name"`
	DealerAddress    string `json:"dealer_address"`
	OwnerName        string `json:"owner_name"`
	OwnerAddress     string `json:"owner_address"`
	MachineModel     string `json:"machine_model"`
	SerialNumber     string `json:"serial_number"`
	InstallDate      string `json:"install_date"`
	InvoiceNumber    string `json:"invoice_number"`
	CompleteSupply   bool   `json:"complete_supply"`
	PdiComplete      bool   `json:"pdi_complete"`
	PtoCorrect       bool   `json:"pto_correct"`
	MachineTestRun   bool   `json:"machine_test_run"`
	SafetyInduction  bool   `json:"safety_induction"`
	OperatorHandbook bool   `json:"operator_handbook"`
	Date             string `json:"date"`
	CompletedBy      string `json:"completed_by"`
}

type LineItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}
