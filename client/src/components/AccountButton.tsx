import styles from '../styles/Header.module.css';

import { Link } from 'react-router-dom';
import { Fragment } from 'react';
import { useUserStore } from '../lib/store';
import LogoutButton from '../forms/LogoutButton';
import { useGetResource } from '../hooks/genericHooks';
import { LineItem } from '../types/miscTypes';

const AccountButton = () => {
    const { isAuthenticated, isAdmin } = useUserStore();
    const { data: lineItem } = useGetResource<LineItem[]>('lineitems');

    return (
        <li className={styles.navItem}>
            {isAuthenticated ? (
                <li className={styles.navListItem}>Account</li>
            ) : (
                <Link to="/login" className={styles.navListItem}>
                    Login
                </Link>
            )}

            <ul className={styles.navDrop}>
                <Fragment>
                    {isAdmin ? (
                        <Fragment>
                            <li className={styles.navDropItem}>
                                <Link to={'/users'}>Users</Link>
                            </li>
                            <li className={styles.navDropItem}>
                                <Link to={'/carousels'}>Carousels</Link>
                            </li>
                            <li className={styles.navDropItem}>
                                <Link to={'/line-items'}>Line Items</Link>
                            </li>
                            <li className={styles.navDropItem}>
                                <Link to={'/warranty'}>Warranty </Link>
                            </li>
                            <li className={styles.navDropItem}>
                                <Link to={'/registrations'}>Registrations</Link>
                            </li>
                            {lineItem && (
                                <li className={styles.navDropItem}>
                                    <Link to={`/checkout/${lineItem[0]?.id}`}>Checkout</Link>
                                </li>
                            )}
                            <LogoutButton mode="listItem" />
                        </Fragment>
                    ) : isAuthenticated ? (
                        <Fragment>
                            <li className={styles.navDropItem}>
                                <Link to={'/warranty'}>Warranty Claims</Link>
                            </li>
                            <li className={styles.navDropItem}>
                                <Link to={'/registrations'}>Registrations</Link>
                            </li>
                            {lineItem && (
                                <li className={styles.navDropItem}>
                                    <Link to={`/checkout/${lineItem[0]?.id}`}>Checkout</Link>
                                </li>
                            )}
                            <LogoutButton mode="listItem" />
                        </Fragment>
                    ) : null}
                </Fragment>
            </ul>
        </li>
    );
};

export default AccountButton;
