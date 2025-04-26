import {defineConfig} from 'vite';
import react from '@vitejs/plugin-react-swc';
import Sitemap from 'vite-plugin-sitemap';

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

const initialRoutes = [...staticRoutes];

const loadDynamicRoutes = async () => {
    try {
        // Fetch all dynamic routes from your dedicated endpoint
        const response = await fetch('https://www.farmec.ie/api/sitemap-data');

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        const {suppliers, spareParts, blogPosts} = data;

        return [
            ...staticRoutes,
            ...suppliers,
            ...spareParts,
            ...blogPosts
        ];
    } catch (error) {
        console.error('Error fetching sitemap data:', error);
        return staticRoutes;
    }
};

let dynamicRoutes = initialRoutes;
loadDynamicRoutes().then(routes => {
    dynamicRoutes = routes;
}).catch(err => {
    console.error('Failed to preload dynamic routes:', err);
});

export default defineConfig({
    plugins: [
        react(),
        Sitemap({
            hostname: 'https://www.farmec.ie',
            generateRobotsTxt: true,
            dynamicRoutes: dynamicRoutes,
            exclude: [
                '/api/*',
                '/login',
                '/line-items',
                '/carousels',
                '/users',
                '/checkout/**',
                '/return',
                '/warranty/**',
                '/registrations/**',
                '/registration/**',
                '/registrations',
                '/registrations',
                '/warranty',
            ]
        }),
    ],
});