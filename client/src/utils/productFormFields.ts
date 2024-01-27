import { useGetResource } from '../hooks/genericHooks';
import { Product, Machine } from '../types/supplierTypes';

export const getFormFields = async (product: Product) => {
    const machines = useGetResource<Machine[]>('machines');

    const machineOptions = machines.data?.map(supplier => ({
        label: supplier.name,
        value: supplier.id,
    }));

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
