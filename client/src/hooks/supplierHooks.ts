import config from '../utils/env';
import { Machine, Supplier } from '../types/supplierTypes';

import { Video } from '../types/videoTypes';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

export const useSuppliers = () => {
    const suppliers = useQuery<Supplier[], Error>({
        queryKey: ['suppliers'],
        queryFn: async () => {
            const response = await fetch(`${config.baseUrl}/api/suppliers`);

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

export const useCreateSupplier = () => {
    const queryClient = useQueryClient();

    const mutateSupplier = useMutation({
        mutationFn: async (supplier: Supplier) => {
            const response = await fetch(`${config.baseUrl}/api/suppliers`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(supplier),
            });

            if (!response.ok) {
                console.log('response:', response);
                throw new Error('Network response was not ok');
            }
            return response.json();
        },

        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['suppliers'] });
        },
    });

    return mutateSupplier;
};

export const useUpdateSupplier = (id: string) => {
    const queryClient = useQueryClient();

    const mutateSupplier = useMutation({
        mutationFn: async (supplier: Supplier) => {
            const response = await fetch(`${config.baseUrl}/api/suppliers/${id}`, {
                method: 'PUT',
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

    return mutateSupplier;
};

export const useDeleteSupplier = (id: string) => {
    const queryClient = useQueryClient();

    const mutateSupplier = useMutation({
        mutationFn: async () => {
            const response = await fetch(`${config.baseUrl}/api/suppliers/${id}`, {
                method: 'DELETE',
            });

            if (!response.ok) {
                throw new Error('network response was not ok');
            }
            return response.json();
        },

        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['suppliers'] });
        },
    });

    return mutateSupplier;
};
