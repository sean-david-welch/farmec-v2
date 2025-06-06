import styles from '../styles/Header.module.css';

import AccountButton from '../components/AccountButton';

import { Link } from 'react-router-dom';
import { Supplier } from '../types/supplierTypes';
import { useLocation } from 'react-router-dom';
import { useGetResource } from '../hooks/genericHooks';
import { useSupplierStore } from '../lib/store';
import { useEffect } from 'react';

const Header: React.FC = () => {
    const location = useLocation();
    const isHomepage = () => location.pathname === '/';

    const { data: suppliers } = useGetResource<Supplier[]>('suppliers');

    const setSuppliers = useSupplierStore(state => state.setSuppliers);

    useEffect(() => {
        if (suppliers) {
            setSuppliers(suppliers);
        }
    }, [suppliers, setSuppliers]);

    return (
        <nav className={isHomepage() ? styles.transparentNav : styles.navbar}>
            <Link to="/" aria-label="logo button">
                <img src="https://static.farmec.ie/farmec_images/farmeclogo.webp" alt="Logo" width="250" height="250" />
            </Link>

            <ul className={styles.navList}>
                <li className={styles.navItem}>
                    <Link to="/about" className={styles.navListItem}>
                        About Us
                    </Link>
                    <ul className={styles.navDrop}>
                        <li className={styles.navDropItem}>
                            <Link to="/about">Staff & Management</Link>
                        </li>
                        <li className={styles.navDropItem}>
                            <a href="/about#timeline">Company History</a>
                        </li>
                        <li className={styles.navDropItem}>
                            <Link to="/about/policies">Terms of Use</Link>
                        </li>
                    </ul>
                </li>

                <li className={styles.navItem}>
                    <Link to="/suppliers" className={styles.navListItem}>
                        Suppliers
                    </Link>
                    {suppliers ? (
                        <ul className={styles.navDrop}>
                            {suppliers.map(supplier => (
                                <li className={styles.navDropItem} key={supplier.id}>
                                    <Link to={`/suppliers/${supplier.id}`}>{supplier.name}</Link>
                                </li>
                            ))}
                        </ul>
                    ) : null}
                </li>

                <li className={styles.navItem}>
                    <Link to="/spareparts" className={styles.navListItem}>
                        Spare Parts
                    </Link>
                    {suppliers ? (
                        <ul className={styles.navDrop}>
                            {suppliers.map(supplier => (
                                <li className={styles.navDropItem} key={supplier.id}>
                                    <Link to={`/spareparts/${supplier.id}`}>{supplier.name}</Link>
                                </li>
                            ))}
                        </ul>
                    ) : null}
                </li>

                <li className={styles.navItem}>
                    <Link to="/blogs" className={styles.navListItem}>
                        Blog
                    </Link>
                    <ul className={styles.navDrop}>
                        <li className={styles.navDropItem}>
                            <Link to="/blogs">Latest Posts</Link>
                        </li>
                        <li className={styles.navDropItem}>
                            <Link to="/blog/exhibitions">Exhibition Information</Link>
                        </li>
                    </ul>
                </li>

                <li className={styles.navItem}>
                    <a href="/#contact" className={styles.navListItem}>
                        Contact
                    </a>
                </li>

                <AccountButton />
            </ul>
        </nav>
    );
};

export default Header;
