import utils from '../styles/Utils.module.css';

import config from '../utils/env';

import { signOutUser } from '../utils/auth';
import { useNavigate } from 'react-router-dom';

const LogoutForm = () => {
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            await signOutUser();

            const response = await fetch(`${config.baseUrl}/api/auth/logout`, {
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
