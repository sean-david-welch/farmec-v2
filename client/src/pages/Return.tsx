import utils from '../styles/Utils.module.css';

import config from '../lib/env';

import { Navigate } from 'react-router-dom';
import { useEffect, useState } from 'react';

const Return: React.FC = () => {
    const baseUrl = config.baseUrl;

    const [status, setStatus] = useState(null);
    const [customerEmail, setCustomerEmail] = useState(null);

    useEffect(() => {
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        const sessionId = urlParams.get('session_id');

        fetch(`${baseUrl}/api/checkout/session-status?session_id=${sessionId}`)
            .then((res) => res.json())
            .then((data) => {
                console.log('data', data);
                setStatus(data.status);
                setCustomerEmail(data.customer_email);
            });
    }, [baseUrl]);

    if (status === 'open') {
        return <Navigate to="/checkout" />;
    }

    if (status === 'complete') {
        return (
            <section id="success">
                <div className={utils.loginSection}>
                    <div className={utils.loginForm}>
                        <h1 className={utils.mainHeading}>Payment Complete</h1>
                        <p className={utils.paragraph}>
                            We appreciate your business! A confirmation email will be sent to {customerEmail}.
                            If you have any questions, please email{' '}
                            <a href="mailto:orders@example.com">info@farmec.ie</a>.
                        </p>
                    </div>
                </div>
            </section>
        );
    }

    return null;
};

export default Return;
