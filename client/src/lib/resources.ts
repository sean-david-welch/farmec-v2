import config from '../lib/env';
import { Resources } from '../types/dataTypes';

const resources: Resources = {
    suppliers: {
        endpoint: new URL('api/suppliers', config.baseUrl).toString(),
        queryKey: 'suppliers',
    },
    machines: {
        endpoint: new URL('api/machines', config.baseUrl).toString(),
        queryKey: 'machines',
    },
    videos: {
        endpoint: new URL('api/videos', config.baseUrl).toString(),
        queryKey: 'videos',
    },
    carousels: {
        endpoint: new URL('api/carousels', config.baseUrl).toString(),
        queryKey: 'carousels',
    },
};

export default resources;
