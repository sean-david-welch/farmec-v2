import styles from '../styles/Header.module.css';
import { useUserStore } from '../utils/context';
import { Link } from 'react-router-dom';

const AccountButton = () => {
    const { isAuthenticated } = useUserStore();

    return (
        <li className={styles.navItem}>
            <Link to={isAuthenticated ? '/account' : '/login'} className={styles.navListItem}>
                {isAuthenticated ? 'Account' : 'Login'}
            </Link>
        </li>
    );
};

export default AccountButton;
