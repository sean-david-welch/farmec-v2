import Layout from '../layouts/Layout';
import Loading from '../layouts/Loading';

import { Suspense, lazy, useEffect } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import { Route, Routes, useLocation, Navigate } from 'react-router-dom';

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
const Warranties = lazy(() => import('../pages/Warranties'));
const Registrations = lazy(() => import('../pages/Registrations'));
const LineItems = lazy(() => import('../pages/LineItems'));
const CarouselAdmin = lazy(() => import('../pages/CarouselAdmin'));
const Users = lazy(() => import('../pages/Users'));
const NotFound = lazy(() => import('../layouts/NotFound'));

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
                        <Route path="/about/about" element={<Navigate replace to="/about" />} />
                        <Route path="/about/policies" element={<Policies />} />

                        <Route path="/suppliers" element={<Suppliers />} />
                        <Route path="/suppliers/suppliers" element={<Navigate replace to="/suppliers" />} />
                        <Route path="/suppliers/:id" element={<SupplierDetail />} />
                        <Route path="/machines/:id" element={<MachineDetail />} />

                        <Route path="/spareparts" element={<Spareparts />} />
                        <Route path="/spareparts/:id" element={<SparepartsDetail />} />
                        <Route path="/spareparts/spare-parts" element={<Navigate replace to="/spareparts" />} />

                        <Route path="/blogs" element={<Blogs />} />
                        <Route path="/blogs/blogs" element={<Navigate replace to="/blogs" />} />
                        <Route path="/blogs/:id" element={<BlogDetail />} />
                        <Route path="/blog/exhibitions" element={<Exhibitions />} />

                        <Route path="/warranty/:id" element={<WarrantyDetail />} />
                        <Route path="/registration/:id" element={<RegistrationDetail />} />

                        <Route path="/checkout/:id" element={<Checkout />} />
                        <Route path="/return" element={<Return />} />

                        <Route path="/login" element={<Login />} />
                        <Route path="/warranty" element={<Warranties />} />
                        <Route path="/registrations" element={<Registrations />} />
                        <Route path="/line-items" element={<LineItems />} />
                        <Route path="/carousels" element={<CarouselAdmin />} />
                        <Route path="/users" element={<Users />} />

                        <Route path="*" element={<NotFound />} />
                    </Routes>
                </Layout>
            </Suspense>
        </Router>
    );
};

export default AppRoutes;
