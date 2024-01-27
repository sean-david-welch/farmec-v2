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
};

export default resources;
