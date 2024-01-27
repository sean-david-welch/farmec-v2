import config from '../lib/env';

import { Employee } from '../types/aboutTypes';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

const url = `${config.baseUrl}/api/employees`;

export const useEmployees = () => {
    const employees = useQuery<Employee, Error>({
        queryKey: ['employees'],
        queryFn: async () => {
            const response = await fetch(url);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    return employees;
};

export const useCreateEmployee = () => {
    const queryClient = useQueryClient();

    const mutateEmployee = useMutation({
        mutationFn: async (employee: Employee) => {
            const response = await fetch(url, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify(employee),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },

        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['employees'] });
        },
    });

    return mutateEmployee;
};

export const useUpdateEmployee = (id: string) => {
    const queryClient = useQueryClient();

    const mutateEmployee = useMutation({
        mutationFn: async (employee: Employee) => {
            const response = await fetch(`${url}/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify(employee),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },

        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['employees'] });
        },
    });

    return mutateEmployee;
};
