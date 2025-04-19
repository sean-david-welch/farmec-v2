import {defineConfig} from 'vite';
import react from '@vitejs/plugin-react-swc';
import Sitemap from 'vite-plugin-sitemap';

const supplierIds = [
    '35607ae9-d92e-4e8f-8dd8-b40654d5eceb',
    'ad309b14-3c19-449c-8dac-78786a95f071',
    'd1eb312f-d00a-44ee-a3e4-45cf9f4783cd',
    'ef0408e4-6e44-4c29-a1d9-e713d4e80d41',
    '67adf8ca-d663-4b16-aa15-4ac9cba0cfc1'
];

const sparePartIds = [
    '35607ae9-d92e-4e8f-8dd8-b40654d5eceb',
    'ad309b14-3c19-449c-8dac-78786a95f071',
    'd1eb312f-d00a-44ee-a3e4-45cf9f4783cd',
    'ef0408e4-6e44-4c29-a1d9-e713d4e80d41',
    '67adf8ca-d663-4b16-aa15-4ac9cba0cfc1'
];

// Basic static routes
const staticRoutes = [
    '/',
    '/about',
    '/about/policies',
    '/suppliers',
    '/spareparts',
    '/blogs',
    '/blog/exhibitions',
    '/return'
];

// Function to generate just the path strings for dynamicRoutes
const generateRoutes = () => {
    return [
        ...staticRoutes,
        ...supplierIds.map(id => `/suppliers/${id}`),
        ...sparePartIds.map(id => `/spareparts/${id}`)
    ];
};

export default defineConfig({
    plugins: [
        react(),
        Sitemap({
            hostname: 'https://www.farmec.ie',
            generateRobotsTxt: true,
            dynamicRoutes: generateRoutes(),
            exclude: [
                '/login',
                '/line-items',
                '/carousels',
                '/users',
                '/checkout/**',
                '/warranty/**',
                '/registration/**',
                '/registrations',
                '/warranty'
            ]
        }),
    ],
});