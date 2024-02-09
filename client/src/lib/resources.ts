import config from '../lib/env';
import { Resources } from '../types/dataTypes';

const resources: Resources = {
    suppliers: {
        endpoint: new URL('api/suppliers', config.baseUrl).toString(),
        queryKey: 'suppliers',
    },
    spareparts: {
        endpoint: new URL('api/spareparts', config.baseUrl).toString(),
        queryKey: 'spareparts',
    },
    machines: {
        endpoint: new URL('api/machines', config.baseUrl).toString(),
        queryKey: 'machines',
    },
    products: {
        endpoint: new URL('api/products', config.baseUrl).toString(),
        queryKey: 'products',
    },
    videos: {
        endpoint: new URL('api/videos', config.baseUrl).toString(),
        queryKey: 'videos',
    },
    blogs: {
        endpoint: new URL('api/blogs', config.baseUrl).toString(),
        queryKey: 'blogs',
    },
    exhibitions: {
        endpoint: new URL('api/exhibitions', config.baseUrl).toString(),
        queryKey: 'exhibitions',
    },
    employees: {
        endpoint: new URL('api/employees', config.baseUrl).toString(),
        queryKey: 'employees',
    },
    timelines: {
        endpoint: new URL('api/timeline', config.baseUrl).toString(),
        queryKey: 'timelines',
    },
    terms: {
        endpoint: new URL('api/terms', config.baseUrl).toString(),
        queryKey: 'terms',
    },
    privacys: {
        endpoint: new URL('api/privacy', config.baseUrl).toString(),
        queryKey: 'privacys',
    },
    lineitems: {
        endpoint: new URL('api/lineitems', config.baseUrl).toString(),
        queryKey: 'lineitems',
    },
    carousels: {
        endpoint: new URL('api/carousels', config.baseUrl).toString(),
        queryKey: 'carousels',
    },
    registrations: {
        endpoint: new URL('api/registrations', config.baseUrl).toString(),
        queryKey: 'registrations',
    },
    warranty: {
        endpoint: new URL('api/warranty', config.baseUrl).toString(),
        queryKey: 'warranty',
    },
    supplierMachine: {
        endpoint: new URL('api/machines/suppliers', config.baseUrl).toString(),
        queryKey: 'machines',
    },
    users: {
        endpoint: new URL('api/auth/users', config.baseUrl).toString(),
        queryKey: 'users',
    },
};

export default resources;
