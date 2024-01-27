import utils from '../styles/Utils.module.css';

import config from '../utils/env';

import { signOutUser } from '../utils/auth';
import { useNavigate } from 'react-router-dom';

const LogoutForm = () => {
    const navigate = useNavigate();

    const url = new URL('api/auth/logout', config.baseUrl);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
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

    return (
        <form onSubmit={handleSubmit} className={utils.form}>
            <button type="submit">Logout</button>
        </form>
    );
};

export default LogoutForm;
