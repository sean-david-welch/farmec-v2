import Layout from '../layouts/Layout';

import { Suspense, lazy, useEffect } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';

import { Route, Routes, useLocation } from 'react-router-dom';
import Loading from '../layouts/Loading';

const Home = lazy(() => import('../pages/Home'));
const About = lazy(() => import('../pages/About'));
const Suppliers = lazy(() => import('../pages/Suppliers'));
const SupplierDetail = lazy(() => import('../pages/details/SupplierDetail'));
const MachineDetail = lazy(() => import('../pages/details/MachineDetail'));
const Spareparts = lazy(() => import('../pages/Spareparts'));
const SparepartsDetail = lazy(() => import('../pages/details/SparepartsDetail'));
const Blogs = lazy(() => import('../pages/Blogs'));
const BlogDetail = lazy(() => import('../pages/details/BlogDetail'));

const Login = lazy(() => import('../pages/Login'));
const Account = lazy(() => import('../pages/Account'));

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
            <Suspense fallback={<Loading />}>
                <Layout>
                    <ScrollToTopPage />
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/about" element={<About />} />
                        <Route path="/suppliers" element={<Suppliers />} />
                        <Route path="/suppliers/:id" element={<SupplierDetail />} />
                        <Route path="/machines/:id" element={<MachineDetail />} />
                        <Route path="/spareparts" element={<Spareparts />} />
                        <Route path="/spareparts/:id" element={<SparepartsDetail />} />

                        <Route path="/blogs" element={<Blogs />} />
                        <Route path="/blogs/:id" element={<BlogDetail />} />

                        <Route path="/login" element={<Login />} />
                        <Route path="/account" element={<Account />} />
                    </Routes>
                </Layout>
            </Suspense>
        </Router>
    );
};

export default AppRoutes;
