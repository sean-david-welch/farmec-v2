import { Video } from '../types/videoTypes';
import { Supplier } from '../types/supplierTypes';

export const getFormFields = (suppliers: Supplier[], video?: Video) => {
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
            defaultValue: video?.supplier_id,
        },
        {
            name: 'web_url',
            label: 'YouTube URL',
            type: 'text',
            placeholder: 'Enter YouTube URL',
            defaultValue: video?.web_url,
        },
    ];
};
