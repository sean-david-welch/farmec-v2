import config from '../utils/env';

import { Employee } from '../types/aboutTypes';
import { useQuery } from '@tanstack/react-query';

export const useEmployees = () => {
    const employees = useQuery<Employee, Error>({
        queryKey: ['employees'],
        queryFn: async () => {
            const response = await fetch(`${config.baseUrl}/api/employees`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        },
    });

    return employees;
};

export const useCreateEmployee = () => {};
