import utils from '../../styles/Utils.module.css';

import config from '../utils/env';

import { signOutUser } from '../utils/auth';
import { useState, useEffect } from 'react';

const LogoutForm = () => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        if (typeof window !== 'undefined') {
            const storedUserData = localStorage.getItem('user');
            if (storedUserData) {
                setUser(JSON.parse(storedUserData));
            }
        }
    }, []);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            await signOutUser();

            const response = await fetch(`${config.baseUrl}/api/auth/logout`);

            if (response.ok) {
                setUser(null);
                localStorage.removeItem('user');
                window.location.href = '/';
            }
        } catch (error) {
            console.error('Error submitting form:', error);
        }
    };

    return user ? (
        <form onSubmit={handleSubmit} className={utils.form}>
            <button type="submit">Logout</button>
        </form>
    ) : (
        <a href="/login">
            <button>Login</button>
        </a>
    );
};

export default LogoutForm;
