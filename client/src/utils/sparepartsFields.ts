import { Sparepart, Supplier } from '../types/supplierTypes';

export const getFormFields = (suppliers: Supplier[], sparepart?: Sparepart) => {
    const supplierOptions = Array.isArray(suppliers)
        ? suppliers.map(supplier => ({
              label: supplier.name,
              value: supplier.id,
          }))
        : [];

    return [
        {
            name: 'supplier_id',
            label: 'Supplier',
            type: 'select',
            options: supplierOptions,
            placeholder: 'Select supplier',
            defaultValue: sparepart?.supplier_id,
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
            label: 'Parts Image ',
            type: 'file',
            placeholder: 'Upload parts image',
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
