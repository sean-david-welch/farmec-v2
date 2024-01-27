import config from '../utils/env';

import { Resources } from '../types/dataTypes';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

const resources: Resources = {
    supplier: {
        endpoint: new URL('api/suppliers', config.baseUrl),
        queryKey: 'suppliers',
    },
};

export const useGetResource = <T>(resourceKey: string) => {
    const { endpoint, queryKey } = resources[resourceKey];

    const resource = useQuery<T, Error>({
        queryKey: [queryKey],
        queryFn: async () => {
            const response = await fetch(endpoint);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    return resource;
};

export const useMutateResource = <T>(resourceKey: string, id?: string) => {
    const queryClient = useQueryClient();
    const { endpoint, queryKey } = resources[resourceKey];

    const buildEndpointUrl = (id?: string) => {
        return id ? new URL(`/${id}`, endpoint).toString() : endpoint;
    };

    const mutate = useMutation({
        mutationFn: async (data: T) => {
            const url = buildEndpointUrl(id);
            const method = id ? 'PUT' : 'POST';

            const response = await fetch(url, {
                method: method,
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify(data),
            });

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(`Network response was not ok: ${response.status} ${errorText}`);
            }
            return response.json();
        },

        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: [queryKey] });
        },
    });

    return mutate;
};

export const useDeleteResource = (resourceKey: string, id: string) => {
    const { endpoint, queryKey } = resources[resourceKey];

    const queryClient = useQueryClient();
    const url = new URL(endpoint + `/${id}`, config.baseUrl).toString();

    const mutateResouce = useMutation({
        mutationFn: async () => {
            const response = await fetch(url, {
                method: 'DELETE',
                credentials: 'include',
            });

            if (!response.ok) {
                throw new Error('network response was not ok');
            }
            return response.json();
        },

        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: [queryKey] });
        },
    });

    return mutateResouce;
};
