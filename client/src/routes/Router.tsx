import { lazy, useEffect } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';

import { Route, Routes, useLocation } from 'react-router-dom';
import Layout from '../layouts/Layout';

const Home = lazy(() => import('../pages/Home'));
const About = lazy(() => import('../pages/About'));
const Suppliers = lazy(() => import('../pages/Suppliers'));
const SupplierDetail = lazy(() => import('../pages/SupplierDetail'));
const Spareparts = lazy(() => import('../pages/Spareparts'));
const Blogs = lazy(() => import('../pages/Blogs'));

const AppRoutes = () => {
    const ScrollToTopPage = () => {
        const { pathname } = useLocation();

        useEffect(() => {
            window.scrollTo(0, 0);
        }, [pathname]);

        return null;
    };
    return (
        <Router>
            <Layout>
                <ScrollToTopPage />
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/about" element={<About />} />
                    <Route path="/suppliers" element={<Suppliers />} />
                    <Route path="/suppliers/:id" element={<SupplierDetail />} />
                    <Route path="/spareparts" element={<Spareparts />} />
                    <Route path="/blogs" element={<Blogs />} />
                </Routes>
            </Layout>
        </Router>
    );
};

export default AppRoutes;
