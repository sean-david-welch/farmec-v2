import utils from '../styles/Utils.module.css';
import styles from '../styles/Blogs.module.css';

import config from '../lib/env';

import { useState } from 'react';
import { signInUser } from '../lib/auth';
import { useNavigate } from 'react-router-dom';

const LoginForm = () => {
    const navigate = useNavigate();
    const url = new URL('api/auth/login', config.baseUrl);

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            const idToken = await signInUser(email, password);

            const response = await fetch(url, {
                method: 'GET',
                headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${idToken}` },
                credentials: 'include',
            });

            if (response.ok) {
                setEmail('');
                setPassword('');
                navigate('/');
            }
        } catch (error) {
            console.error('Error submitting form:', error);
            setErrorMessage('An unexpected error occurred. Please try again later.');
        }
    };

    return (
        <section id="form">
            <form onSubmit={handleSubmit} className={utils.form}>
                <label>Email:</label>
                <input
                    type="email"
                    value={email}
                    onChange={(e) => {
                        setEmail(e.target.value);
                        setErrorMessage('');
                    }}
                    required
                />
                <label>Password:</label>
                <input
                    type="password"
                    value={password}
                    onChange={(e) => {
                        setPassword(e.target.value);
                        setErrorMessage('');
                    }}
                    required
                />
                {errorMessage && <div className={styles.errorMessage}>{errorMessage}</div>}
                <button type="submit">Login</button>
            </form>
        </section>
    );
};

export default LoginForm;
