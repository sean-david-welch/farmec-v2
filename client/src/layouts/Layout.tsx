// src/components/Layout.jsx
import { Component, JSX } from 'solid-js';
import Header from './Header';
import Footer from './Footer';

interface LayoutProps {
    children: JSX.Element;
}

const Layout: Component<LayoutProps> = props => {
    return (
        <div>
            <Header />
            <main>{props.children}</main>
            <Footer />
        </div>
    );
};

export default Layout;
