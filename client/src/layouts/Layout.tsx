import Header from './Header';
import Footer from './Footer';

interface LayoutProps {
    children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
    return (
        <div>
            <Header />
            <main className="min-h-screen max-w-full overflow-x-hidden mx-auto">{children}</main>
            <Footer />
        </div>
    );
};

export default Layout;
