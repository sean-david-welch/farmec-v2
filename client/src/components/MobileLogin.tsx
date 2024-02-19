import styles from '../styles/Sidebar.module.css';

import LogoutButton from '../forms/LogoutButton';

import { Link } from 'react-router-dom';
import { Fragment } from 'react';
import { LineItem } from '../types/miscTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
    faUserCircle,
    faUsers,
    faImage,
    faPenToSquare,
    faCartShopping,
    faSitemap,
} from '@fortawesome/free-solid-svg-icons';

interface Props {
    onClick: () => void;
}

const MobileButton: React.FC<Props> = ({ onClick }) => {
    const { isAuthenticated, isAdmin } = useUserStore();
    const { data: lineItem } = useGetResource<LineItem[]>('lineitems');

    return (
        <Fragment>
            {isAuthenticated ? (
                <Fragment>
                    {isAdmin ? (
                        <Fragment>
                            <Link className={styles.navItem} to={'/users'} onClick={onClick}>
                                <FontAwesomeIcon icon={faUsers} />
                                Users
                            </Link>

                            <Link className={styles.navItem} to={'/carousels'} onClick={onClick}>
                                <FontAwesomeIcon icon={faImage} />
                                Carousels
                            </Link>

                            <Link className={styles.navItem} to={'/line-items'} onClick={onClick}>
                                <FontAwesomeIcon icon={faSitemap} />
                                Line Items
                            </Link>

                            <Link className={styles.navItem} to={'/warranty'} onClick={onClick}>
                                <FontAwesomeIcon icon={faPenToSquare} />
                                Warranty
                            </Link>

                            <Link className={styles.navItem} to={'/registrations'} onClick={onClick}>
                                <FontAwesomeIcon icon={faPenToSquare} />
                                Registrations
                            </Link>

                            {lineItem && (
                                <Link className={styles.navItem} to={`/checkout/${lineItem[0]?.id}`} onClick={onClick}>
                                    <FontAwesomeIcon icon={faCartShopping} />
                                    Checkout
                                </Link>
                            )}
                            <LogoutButton mode="button" />
                        </Fragment>
                    ) : isAuthenticated ? (
                        <Fragment>
                            <Link className={styles.navItem} to={'/warranty'} onClick={onClick}>
                                <FontAwesomeIcon icon={faPenToSquare} />
                                Warranty Claims
                            </Link>

                            <Link className={styles.navItem} to={'/registrations'} onClick={onClick}>
                                <FontAwesomeIcon icon={faPenToSquare} />
                                Registrations
                            </Link>

                            {lineItem && (
                                <Link className={styles.navItem} to={`/checkout/${lineItem[0]?.id}`} onClick={onClick}>
                                    <FontAwesomeIcon icon={faCartShopping} />
                                    Checkout
                                </Link>
                            )}
                            <LogoutButton mode="button" />
                        </Fragment>
                    ) : null}
                </Fragment>
            ) : (
                <Link to={'/login'} className={styles.navItem} onClick={onClick}>
                    <FontAwesomeIcon icon={faUserCircle} />
                    Login
                </Link>
            )}
        </Fragment>
    );
};

export default MobileButton;
