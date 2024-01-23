import { lazy } from 'solid-js';
import { Router, Route } from '@solidjs/router';

const Home = lazy(() => import('./routes/Home'));
const About = lazy(() => import('./routes/About'));
const Suppliers = lazy(() => import('./routes/Suppliers'));
const SupplierDetail = lazy(() => import('./routes/SupplierDetail'));
const Spareparts = lazy(() => import('./routes/Spareparts'));
const Blogs = lazy(() => import('./routes/Blogs'));

function AppRoutes() {
    return (
        <Router>
            <Route path="/" component={Home} />
            <Route path="/about" component={About} />
            <Route path="/suppliers" component={Suppliers} />
            <Route path="/suppliers/:id" component={SupplierDetail} />
            <Route path="/spareparts" component={Spareparts} />
            <Route path="/blogs" component={Blogs} />
        </Router>
    );
}

export default AppRoutes;
