import { Machine, Product } from '../types/supplierTypes';

import { useGetResourceById } from './genericHooks';

export const useMachineDetails = (id: string) => {
    const machine = useGetResourceById<Machine>('machines', id);
    const products = useGetResourceById<Product[]>('products', id);

    return { machine, products };
};
