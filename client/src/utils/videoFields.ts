import { Video } from '../types/videoTypes';
import { Supplier } from '../types/supplierTypes';
import { useGetResource } from '../hooks/genericHooks';

export const getFormFields = async (video: Video) => {
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
            defaultValue: video?.supplierId,
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
