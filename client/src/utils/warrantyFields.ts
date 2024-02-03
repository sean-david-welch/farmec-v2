import { WarrantyClaim, PartsRequired } from '../types/miscTypes';

export const getWarrantyFields = (warrantyClaim?: WarrantyClaim) => [
    {
        name: 'dealer',
        label: 'Dealer',
        type: 'text',
        placeholder: 'Enter dealer',
        defaultValue: warrantyClaim?.dealer,
    },
    {
        name: 'dealer_contact',
        label: 'Dealer Contact',
        type: 'text',
        placeholder: 'Enter dealer contact',
        defaultValue: warrantyClaim?.dealer_contact,
    },
    {
        name: 'owner_name',
        label: 'Owner Name',
        type: 'text',
        placeholder: 'Enter owner name',
        defaultValue: warrantyClaim?.owner_name,
    },
    {
        name: 'machine_model',
        label: 'Machine Model',
        type: 'text',
        placeholder: 'Enter machine model',
        defaultValue: warrantyClaim?.machine_model,
    },
    {
        name: 'serial_number',
        label: 'Serial Number',
        type: 'text',
        placeholder: 'Enter serial number',
        defaultValue: warrantyClaim?.serial_number,
    },
    {
        name: 'install_date',
        label: 'Install Date',
        type: 'text',
        placeholder: 'Enter install date',
        defaultValue: warrantyClaim?.install_date,
    },
    {
        name: 'failure_date',
        label: 'Failure Date',
        type: 'text',
        placeholder: 'Enter failure date',
        defaultValue: warrantyClaim?.failure_date,
    },
    {
        name: 'repair_date',
        label: 'Repair Date',
        type: 'text',
        placeholder: 'Enter repair date',
        defaultValue: warrantyClaim?.repair_date,
    },
    {
        name: 'failure_details',
        label: 'Failure Details',
        type: 'text',
        placeholder: 'Enter failure details',
        defaultValue: warrantyClaim?.failure_details,
    },
    {
        name: 'repair_details',
        label: 'Repair Details',
        type: 'text',
        placeholder: 'Enter repair details',
        defaultValue: warrantyClaim?.repair_details,
    },
    {
        name: 'labour_hours',
        label: 'Labour Hours',
        type: 'text',
        placeholder: 'Enter labour hours',
        defaultValue: warrantyClaim?.labour_hours,
    },
    {
        name: 'completed_by',
        label: 'Completed By',
        type: 'text',
        placeholder: 'Enter completed by',
        defaultValue: warrantyClaim?.completed_by,
    },
];

export const getPartFields = (partsRequired: Partial<PartsRequired> = {}, index?: number) => [
    {
        name: `part_number_${index}`,
        label: 'Part Number',
        type: 'text',
        placeholder: 'Enter part number',
        defaultValue: partsRequired?.part_number,
    },
    {
        name: `quantity_needed_${index}`,
        label: 'Quantity Needed',
        type: 'text',
        placeholder: 'Enter quantity needed',
        defaultValue: partsRequired?.quantity_needed,
    },
    {
        name: `invoice_number_${index}`,
        label: 'Invoice Number',
        type: 'text',
        placeholder: 'Enter invoice number',
        defaultValue: partsRequired?.invoice_number,
    },
    {
        name: `part_description_${index}`,
        label: 'Part Description',
        type: 'text',
        placeholder: 'Enter  description',
        defaultValue: partsRequired?.description,
    },
];
