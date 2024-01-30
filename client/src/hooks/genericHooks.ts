import config from '../lib/env';
import resources from '../lib/resources';

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

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

export const useGetResourceById = <T>(resourceKey: string, id: string) => {
    const { endpoint, queryKey } = resources[resourceKey];

    const url = `${endpoint}/${id}`;

    const resource = useQuery<T, Error>({
        queryKey: [queryKey],
        queryFn: async () => {
            const response = await fetch(url);

            if (!response.ok) throw new Error('Network response was not ok');
            return response.json();
        },
    });

    return resource;
};

export const useMutateResource = <T>(resourceKey: string, id?: string) => {
    const queryClient = useQueryClient();
    const { endpoint, queryKey } = resources[resourceKey];

    const buildEndpointUrl = (id?: string) => {
        return id ? `${endpoint}/${id}` : endpoint;
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

export const useCreateOrUpdateResource = <T>(resourceKey: string, id?: string) => {
    const {
        mutateAsync: createResource,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<T>(resourceKey);

    const {
        mutateAsync: updateResource,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<T>(resourceKey, id);

    return {
        createResource,
        updateResource,
        isCreateError,
        isUpdateError,
        createError,
        updateError,
    };
};
