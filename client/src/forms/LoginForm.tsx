import utils from '../styles/Utils.module.css';
import styles from '../styles/Blogs.module.css';

import config from '../utils/env';

import { useState } from 'react';
import { signInUser } from '../utils/auth';

const LoginForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            const idToken = await signInUser(email, password);
            console.log('IdToken', idToken);

            const response = await fetch(`${config.baseUrl}/api/auth/login`, {
                method: 'GET',
                headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${idToken}` },
            }).then(response => response.json());

            if (response.ok) {
                setEmail('');
                setPassword('');
                window.location.href = '/';
            }
        } catch (error: any) {
            if (error.response && error.response.status === 401) {
                setErrorMessage('Incorrect email or password.');
            } else {
                setErrorMessage('An unexpected error occurred. Please try again later.');
            }

            console.error('Error submitting form:', error.message || error);
        }
    };

    return (
        <section id="form">
            <form onSubmit={handleSubmit} className={utils.form}>
                <label>Email:</label>
                <input
                    type="email"
                    value={email}
                    onChange={e => {
                        setEmail(e.target.value);
                        setErrorMessage('');
                    }}
                    required
                />
                <label>Password:</label>
                <input
                    type="password"
                    value={password}
                    onChange={e => {
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
