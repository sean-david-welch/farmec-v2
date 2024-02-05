import config from '../lib/env';

import { useEffect, useState } from 'react';
import { loadStripe } from '@stripe/stripe-js';
import { EmbeddedCheckoutProvider, EmbeddedCheckout } from '@stripe/react-stripe-js';
import { useParams } from 'react-router-dom';

const stripePromise = loadStripe(config.stripePublicKeyTest);

const CheckoutForm: React.FC = () => {
    const id = useParams<{ id: string }>().id;
    const baseUrl = config.baseUrl;
    const [clientSecret, setClientSecret] = useState('');

    useEffect(() => {
        fetch(`${baseUrl}/api/checkout/create-checkout-session/${id}`, {
            method: 'POST',
        })
            .then((res) => res.json())
            .then((data) => {
                setClientSecret(data.session.client_secret);
            })
            .catch((error) => {
                console.error('Error fetching checkout session:', error);
            });
    }, [id, baseUrl]);

    return (
        <div id="checkout">
            {clientSecret && (
                <EmbeddedCheckoutProvider stripe={stripePromise} options={{ clientSecret }}>
                    <EmbeddedCheckout />
                </EmbeddedCheckoutProvider>
            )}
        </div>
    );
};

export default CheckoutForm;
