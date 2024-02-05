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
                <p>
                    We appreciate your business! A confirmation email will be sent to {customerEmail}. If you
                    have any questions, please email{' '}
                    <a href="mailto:orders@example.com">orders@example.com</a>.
                </p>
            </section>
        );
    }

    return null;
};

export default Return;
