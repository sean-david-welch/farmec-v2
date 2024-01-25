import Supplier from '../types/supplier';
import Machine from '../types/machine';

import { Video } from '../types/video';
import { useQuery } from '@tanstack/react-query';

const fetchSupplier = async (id: string): Promise<Supplier> => {
    const response = await fetch(`http://localhost:8080/api/suppliers/${id}`);

    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
};

const fetchMachines = async (id: string): Promise<Machine[]> => {
    const response = await fetch(`http://localhost:8080/api/machines/${id}`);
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
};

const fetchVideos = async (id: string): Promise<Video[]> => {
    const response = await fetch(`http://localhost:8080/api/videos/${id}`);
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
};

export const useSupplierDetails = (id: string) => {
    const supplierQueryKey = ['supplier', id];
    const machinesQueryKey = ['machines', id];
    const videosQueryKey = ['videos', id];

    const supplier = useQuery<Supplier, Error>({
        queryKey: supplierQueryKey,
        queryFn: () => fetchSupplier(id),
    });

    const machines = useQuery<Machine[], Error>({
        queryKey: machinesQueryKey,
        queryFn: () => fetchMachines(id),
    });

    const videos = useQuery<Video[], Error>({
        queryKey: videosQueryKey,
        queryFn: () => fetchVideos(id),
    });

    return { supplier, machines, videos };
};
