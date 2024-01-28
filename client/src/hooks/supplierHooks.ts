import { Video } from '../types/videoTypes';
import { Machine, Product, Supplier } from '../types/supplierTypes';

import { useGetResourceById } from './genericHooks';

export const useSupplierDetails = (id: string) => {
    const supplier = useGetResourceById<Supplier>('suppliers', id);
    const machines = useGetResourceById<Machine[]>('machines', id);
    const videos = useGetResourceById<Video[]>('videos', id);

    return { supplier, machines, videos };
};

export const useMachineDetails = (id: string) => {
    const machine = useGetResourceById<Machine>('machines', id);
    const products = machine.data?.id ? useGetResourceById<Product[]>('products', machine.data.id) : null;

    return { machine, products };
};
