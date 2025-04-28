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
        let response = await fetch("https://www.farmec.ie/api/sitemap-data");
        if (response.status === 404) {
            console.log("Production API not found, falling back to localhost");
            response = await fetch("http://localhost:8080/api/sitemap-data");
        }
        if (!response.ok) {
            console.error(`HTTP error! Status: ${response.status}`);
            return staticRoutes;
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

// let dynamicRoutes = initialRoutes;
// loadDynamicRoutes().then(routes => {
//     dynamicRoutes = routes;
// }).catch(err => {
//     console.error('Failed to preload dynamic routes:', err);
// });

export default defineConfig({
    plugins: [
        react(),
        Sitemap({
            hostname: 'https://www.farmec.ie',
            generateRobotsTxt: true,
            // dynamicRoutes: dynamicRoutes,
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