import Supplier from '../types/supplier';

import { useQuery } from '@tanstack/react-query';

const fetchSuppliers = async (): Promise<Supplier[]> => {
    const response = await fetch(`http://localhost:8080/api/suppliers`);

    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
};

export const useSuppliers = () => {
    const suppliersQueryKey = ['suppliers'];

    const suppliers = useQuery<Supplier[], Error>({
        queryKey: suppliersQueryKey,
        queryFn: () => fetchSuppliers(),
    });

    return { suppliers };
};
