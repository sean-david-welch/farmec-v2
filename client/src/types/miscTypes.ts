export interface Carousel {
    id?: string;
    name: string;
    image: string;
}

export interface DealerOwnerInfo {
    dealer: string;
    ownerName: string;
}

export interface WarrantyClaim {
    id?: string;
    dealer: string;
    dealer_contact: string;
    owner_name: string;
    machine_model: string;
    serial_number: string;
    install_date: string;
    failure_date: string;
    repair_date: string;
    failure_details: string;
    repair_details: string;
    labour_hours: string;
    completed_by: string;
    created?: string;
}

export interface PartsRequired {
    id?: string;
    warranty_id?: string;
    part_number: string;
    quantity_needed: string;
    invoice_number: string;
    description: string;
}

export interface WarrantyParts {
    warranty: WarrantyClaim;
    parts: PartsRequired[];
}

export interface MachineRegistration {
    id?: string;
    dealer_name: string;
    dealer_address: string;
    owner_name: string;
    owner_address: string;
    machine_model: string;
    serial_number: string;
    install_date: string;
    invoice_number: string;
    complete_supply: boolean;
    pdi_complete: boolean;
    pto_correct: boolean;
    machine_test_run: boolean;
    safety_induction: boolean;
    operator_handbook: boolean;
    date: string;
    completed_by: string;
    created?: string;
}

export interface LineItem {
    id?: string;
    name: string;
    price: number;
    image: string;
}
