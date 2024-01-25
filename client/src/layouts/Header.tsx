import styles from '../styles/Header.module.css';

import AccountButton from '../components/AccountButton';

import { useLocation } from 'react-router-dom';
import { useSuppliers } from '../hooks/useSuppliers';

const Header: React.FC = () => {
    const location = useLocation();
    const isHomepage = () => location.pathname === '/';

    const { suppliers } = useSuppliers();

    return (
        <nav className={isHomepage() ? styles.transparentNav : styles.navbar}>
            <a href="/" aria-label="logo button">
                <img
                    src="https://d3eerclezczw8.cloudfront.net/farmec_images/farmeclogo.webp"
                    alt="Logo"
                    width="250"
                    height="250"
                />
            </a>

            <ul className={styles.navList}>
                <li className={styles.navItem}>
                    <a href="/about" className={styles.navListItem}>
                        About Us
                    </a>
                    <ul className={styles.navDrop}>
                        <li className={styles.navDropItem}>
                            <a href="/about">Staff & Management</a>
                        </li>
                        <li className={styles.navDropItem}>
                            <a href="/about#timeline">Company History</a>
                        </li>
                        <li className={styles.navDropItem}>
                            <a href="/about/policies">Terms of Use</a>
                        </li>
                    </ul>
                </li>

                <li className={styles.navItem}>
                    <a href="/suppliers" className={styles.navListItem}>
                        Suppliers
                    </a>
                    {suppliers.data && (
                        <ul className={styles.navDrop}>
                            {suppliers.data.map(supplier => (
                                <li className={styles.navDropItem} key={supplier.id}>
                                    <a href={`/suppliers/${supplier.id}`}>{supplier.name}</a>
                                </li>
                            ))}
                        </ul>
                    )}
                </li>

                <li className={styles.navItem}>
                    <a href="/spareparts" className={styles.navListItem}>
                        Spareparts
                    </a>

                    {suppliers.data && (
                        <ul className={styles.navDrop}>
                            {suppliers.data.map(supplier => (
                                <li className={styles.navDropItem} key={supplier.id}>
                                    <a href={`/spareparts/${supplier.id}`}>{supplier.name}</a>
                                </li>
                            ))}
                        </ul>
                    )}
                </li>

                <li className={styles.navItem}>
                    <a href="/blogs" className={styles.navListItem}>
                        Blog
                    </a>
                    <ul className={styles.navDrop}>
                        <li className={styles.navDropItem}>
                            <a href="/blog">Latest Posts</a>
                        </li>
                        <li className={styles.navDropItem}>
                            <a href="/blog/exhibitions">Exhibition Information</a>
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
