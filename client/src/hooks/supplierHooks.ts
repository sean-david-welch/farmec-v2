import Machine from '../types/machine';
import Supplier from '../types/supplier';

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { Video } from '../types/video';

export const useSuppliers = () => {
    const suppliers = useQuery<Supplier[], Error>({
        queryKey: ['suppliers'],
        queryFn: async () => {
            const response = await fetch(`http://localhost:8080/api/suppliers`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    return { suppliers };
};

export const useSupplierDetails = (id: string) => {
    const supplier = useQuery<Supplier, Error>({
        queryKey: ['supplier', id],
        queryFn: async () => {
            const response = await fetch(`http://localhost:8080/api/suppliers/${id}`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    const machines = useQuery<Machine[], Error>({
        queryKey: ['machines', id],
        queryFn: async () => {
            const response = await fetch(`http://localhost:8080/api/machines/${id}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    const videos = useQuery<Video[], Error>({
        queryKey: ['videos', id],
        queryFn: async () => {
            const response = await fetch(`http://localhost:8080/api/videos/${id}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    return { supplier, machines, videos };
};

export const useCreateSupplier = () => {
    const queryClient = useQueryClient();

    const mutationFn = useMutation({
        mutationFn: async (supplier: Supplier) => {
            const response = await fetch(`http://localhost:8080/api/suppliers`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(supplier),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },

        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['suppliers'] });
        },
    });

    return mutationFn;
};
