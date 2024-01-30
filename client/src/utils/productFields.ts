import { Product, Machine } from '../types/supplierTypes';

export const getFormFields = (machines: Machine[], product?: Product) => {
    const machineOptions = Array.isArray(machines)
        ? machines.map(machine => ({
              label: machine.name,
              value: machine.id,
          }))
        : [];

    return [
        {
            name: 'machine',
            label: 'Macine',
            type: 'select',
            options: machineOptions,
            placeholder: 'Select machine',
            defaultValue: product?.machineId,
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
