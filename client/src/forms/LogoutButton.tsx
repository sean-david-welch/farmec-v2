import utils from '../styles/Utils.module.css';
import styles from '../styles/Header.module.css';

import config from '../lib/env';

import { signOutUser } from '../lib/auth';
import { useNavigate } from 'react-router-dom';

interface Props {
    mode: 'button' | 'listItem';
}

const LogoutButton: React.FC<Props> = ({ mode }) => {
    const navigate = useNavigate();
    const url = new URL('api/auth/logout', config.baseUrl);

    const handleLogout = async (event: React.MouseEvent<HTMLElement>) => {
        event.preventDefault();

        try {
            await signOutUser();

            const response = await fetch(url, {
                credentials: 'include',
            });

            if (response.ok) navigate('/');
        } catch (error) {
            console.error('Error submitting form:', error);
        }
    };

    if (mode === 'button') {
        return (
            <button className={utils.btnForm} onClick={handleLogout}>
                Logout
            </button>
        );
    }

    if (mode === 'listItem') {
        return (
            <button className={styles.navDropItem} onClick={handleLogout}>
                Logout
            </button>
        );
    }

    return null;
};

export default LogoutButton;
