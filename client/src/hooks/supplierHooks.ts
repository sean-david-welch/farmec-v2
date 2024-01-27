import { Video } from '../types/videoTypes';
import { Machine, Supplier } from '../types/supplierTypes';

import { useGetResourceById } from './genericHooks';

export const useSupplierDetails = (id: string) => {
    const supplier = useGetResourceById<Supplier>('suppliers', id);
    const machines = useGetResourceById<Machine[]>('machines', id);
    const videos = useGetResourceById<Video[]>('videos', id);

    return { supplier, machines, videos };
};
