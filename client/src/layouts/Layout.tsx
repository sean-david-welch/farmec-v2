import Header from './Header';
import Footer from './Footer';
import { useLocation } from 'react-router-dom';

interface LayoutProps {
    children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
    const location = useLocation();

    const mainClass = location.pathname === '/' ? 'flex-grow' : 'flex-grow my-32';

    return (
        <div className="mx-auto flex min-h-screen max-w-full flex-col overflow-x-hidden">
            <Header />
            <main className={mainClass}>{children}</main>
            <Footer />
        </div>
    );
};

export default Layout;
