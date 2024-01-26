import { useMutation, useQueryClient } from '@tanstack/react-query';
import config from '../utils/env';

interface ResourceData {
    id: string;
    route: string;
    queryKey: string;
}

export const useDeleteResource = (resourceData: ResourceData) => {
    const { id, route, queryKey } = resourceData;

    const queryClient = useQueryClient();
    const url = `${config.baseUrl}/api/${route}/${id}`;

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
            queryClient.invalidateQueries({ queryKey: [`${queryKey}`] });
        },
    });

    return mutateResouce;
};
