import { Product, Machine } from '../types/supplierTypes';

export const getFormFields = (machine: Machine, product?: Product) => {
    const machineOptions = [
        {
            label: machine.name,
            value: machine.id,
        },
    ];

    return [
        {
            name: 'machine_id',
            label: 'Machine',
            type: 'select',
            options: machineOptions,
            placeholder: 'Select machine',
            defaultValue: product?.machine_id,
        },
        {
            name: 'name',
            label: 'Name',
            type: 'text',
            placeholder: 'Enter name',
            defaultValue: product?.name,
        },
        {
            name: 'product_image',
            label: 'Product Image',
            type: 'file',
            placeholder: 'Upload product image',
        },
        {
            name: 'description',
            label: 'Description',
            type: 'text',
            placeholder: 'Enter description',
            defaultValue: product?.description,
        },
        {
            name: 'product_link',
            label: 'Product Link',
            type: 'text',
            placeholder: 'Enter product link',
            defaultValue: product?.product_link,
        },
    ];
};
