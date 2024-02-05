import Layout from '../layouts/Layout';

import { Suspense, lazy, useEffect } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';

import { Route, Routes, useLocation } from 'react-router-dom';
import Loading from '../layouts/Loading';

const Home = lazy(() => import('../pages/Home'));
const About = lazy(() => import('../pages/About'));
const Policies = lazy(() => import('../pages/Policies'));
const Suppliers = lazy(() => import('../pages/Suppliers'));
const SupplierDetail = lazy(() => import('../pages/details/SupplierDetail'));
const MachineDetail = lazy(() => import('../pages/details/MachineDetail'));
const Spareparts = lazy(() => import('../pages/Spareparts'));
const SparepartsDetail = lazy(() => import('../pages/details/SparepartsDetail'));
const Blogs = lazy(() => import('../pages/Blogs'));
const BlogDetail = lazy(() => import('../pages/details/BlogDetail'));
const Exhibitions = lazy(() => import('../pages/Exhibitions'));
const WarrantyDetail = lazy(() => import('../pages/details/WarrantyDetail'));
const RegistrationDetail = lazy(() => import('../pages/details/RegistrationDetail'));
const Checkout = lazy(() => import('../pages/Checkout'));
const Return = lazy(() => import('../pages/Return'));
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
                        <Route path="/about/policies" element={<Policies />} />

                        <Route path="/suppliers" element={<Suppliers />} />
                        <Route path="/suppliers/:id" element={<SupplierDetail />} />
                        <Route path="/machines/:id" element={<MachineDetail />} />

                        <Route path="/spareparts" element={<Spareparts />} />
                        <Route path="/spareparts/:id" element={<SparepartsDetail />} />

                        <Route path="/blogs" element={<Blogs />} />
                        <Route path="/blogs/:id" element={<BlogDetail />} />
                        <Route path="/blog/exhibitions" element={<Exhibitions />} />

                        <Route path="/warranty/:id" element={<WarrantyDetail />} />
                        <Route path="/registration/:id" element={<RegistrationDetail />} />

                        <Route path="/checkout/:id" element={<Checkout />} />
                        <Route path="/return" element={<Return />} />

                        <Route path="/login" element={<Login />} />
                        <Route path="/account" element={<Account />} />
                    </Routes>
                </Layout>
            </Suspense>
        </Router>
    );
};

export default AppRoutes;
