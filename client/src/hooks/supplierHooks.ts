import config from '../utils/env';

import { Machine, Supplier } from '../types/supplierTypes';

import { Video } from '../types/videoTypes';
import { useQuery } from '@tanstack/react-query';

export const useSuppliers = () => {
    const suppliers = useQuery<Supplier[], Error>({
        queryKey: ['suppliers'],
        queryFn: async () => {
            const response = await fetch(`${config.baseUrl}/suppliers`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    return suppliers;
};

export const useSupplierDetails = (id: string) => {
    const supplier = useQuery<Supplier, Error>({
        queryKey: ['supplier', id],
        queryFn: async () => {
            const response = await fetch(`${config.baseUrl}/api/suppliers/${id}`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    const machines = useQuery<Machine[], Error>({
        queryKey: ['machines', id],
        queryFn: async () => {
            const response = await fetch(`${config.baseUrl}/api/machines/${id}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    const videos = useQuery<Video[], Error>({
        queryKey: ['videos', id],
        queryFn: async () => {
            const response = await fetch(`${config.baseUrl}/api/videos/${id}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    return { supplier, machines, videos };
};
