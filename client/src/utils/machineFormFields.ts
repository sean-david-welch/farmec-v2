import { useGetResource } from '../hooks/genericHooks';
import { Machine, Supplier } from '../types/supplierTypes';

export const getFormFields = async (machine: Machine) => {
    const suppliers = useGetResource<Supplier[]>('suppliers');

    const supplierOptions = suppliers.data?.map(supplier => ({
        label: supplier.name,
        value: supplier.id,
    }));

    return [
        {
            name: 'supplier',
            label: 'Supplier',
            type: 'select',
            options: supplierOptions,
            placeholder: 'Select supplier',
            defaultValue: machine?.supplierId,
        },
        {
            name: 'name',
            label: 'Name',
            type: 'text',
            placeholder: 'Enter name',
            defaultValue: machine?.name,
        },
        {
            name: 'machine_image',
            label: 'Machine Image',
            type: 'file',
            placeholder: 'Upload machine image',
        },
        {
            name: 'description',
            label: 'Description',
            type: 'text',
            placeholder: 'Enter description',
            defaultValue: machine?.description,
        },
        {
            name: 'machine_link',
            label: 'Machine Link',
            type: 'text',
            placeholder: 'Enter machine link',
            defaultValue: machine?.machine_link,
        },
    ];
};