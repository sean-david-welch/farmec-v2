import styles from '../styles/Header.module.css';
import { useState, useEffect } from 'react';

const AccountButton = () => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        const storedUserData = localStorage.getItem('user');
        if (storedUserData) {
            const storedUser = JSON.parse(storedUserData);
            setUser(storedUser);
        }
    }, []); // The empty array ensures this effect runs once on mount

    return (
        <li className={styles.navItem}>
            <a href={user ? '/account' : '/login'} className={styles.navListItem}>
                {user ? 'Account' : 'Login'}
            </a>
        </li>
    );
};

export default AccountButton;
