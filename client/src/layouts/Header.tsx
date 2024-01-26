import styles from '../styles/Header.module.css';

import AccountButton from '../components/AccountButton';

import { useLocation } from 'react-router-dom';
import { useSuppliers } from '../hooks/supplierHooks';

import { Link } from 'react-router-dom';

const Header: React.FC = () => {
    const location = useLocation();
    const isHomepage = () => location.pathname === '/';

    const suppliers = useSuppliers();
    return (
        <nav className={isHomepage() ? styles.transparentNav : styles.navbar}>
            <Link to="/" aria-label="logo button">
                <img
                    src="https://d3eerclezczw8.cloudfront.net/farmec_images/farmeclogo.webp"
                    alt="Logo"
                    width="250"
                    height="250"
                />
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
                            <Link to="/about#timeline">Company History</Link>
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
                    {suppliers.data && (
                        <ul className={styles.navDrop}>
                            {suppliers.data.map(supplier => (
                                <li className={styles.navDropItem} key={supplier.id}>
                                    <Link to={`/suppliers/${supplier.id}`}>{supplier.name}</Link>
                                </li>
                            ))}
                        </ul>
                    )}
                </li>

                <li className={styles.navItem}>
                    <Link to="/spareparts" className={styles.navListItem}>
                        Spareparts
                    </Link>

                    {suppliers.data && (
                        <ul className={styles.navDrop}>
                            {suppliers.data.map(supplier => (
                                <li className={styles.navDropItem} key={supplier.id}>
                                    <Link to={`/spareparts/${supplier.id}`}>{supplier.name}</Link>
                                </li>
                            ))}
                        </ul>
                    )}
                </li>

                <li className={styles.navItem}>
                    <Link to="/blogs" className={styles.navListItem}>
                        Blog
                    </Link>
                    <ul className={styles.navDrop}>
                        <li className={styles.navDropItem}>
                            <Link to="/blog">Latest Posts</Link>
                        </li>
                        <li className={styles.navDropItem}>
                            <Link to="/blog/exhibitions">Exhibition Information</Link>
                        </li>
                    </ul>
                </li>

                <li className={styles.navItem}>
                    <Link to="/#contact" className={styles.navListItem}>
                        Contact
                    </Link>
                </li>

                <AccountButton />
            </ul>
        </nav>
    );
};

export default Header;
