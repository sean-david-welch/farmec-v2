import { useGetResource } from '../hooks/genericHooks';
import { Sparepart, Supplier } from '../types/supplierTypes';

export const getFormFields = async (sparepart: Sparepart) => {
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
            defaultValue: sparepart?.supplierId,
        },
        {
            name: 'name',
            label: 'Name',
            type: 'text',
            placeholder: 'Enter name',
            defaultValue: sparepart?.name,
        },
        {
            name: 'parts_image',
            label: 'Parts Image (Max 10MB)',
            type: 'file',
            placeholder: 'Upload parts image',
        },
        {
            name: 'pdf_link',
            label: 'PDF Link (Max 10MB)',
            type: 'file',
            placeholder: 'Enter pdf_link',
        },
        {
            name: 'spare_parts_link',
            label: 'Spare Parts Link',
            type: 'text',
            placeholder: 'Enter sparepart link',
            defaultValue: sparepart?.spare_parts_link,
        },
    ];
};
