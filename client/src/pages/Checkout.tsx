import utils from '../styles/Utils.module.css';

import config from '../lib/env';

import { useEffect, useState } from 'react';
import { loadStripe } from '@stripe/stripe-js';
import { EmbeddedCheckoutProvider, EmbeddedCheckout } from '@stripe/react-stripe-js';
import { useParams } from 'react-router-dom';

const stripePromise = loadStripe(config.stripePublicKeyTest);

const CheckoutForm: React.FC = () => {
    const [clientSecret, setClientSecret] = useState('');

    const id = useParams<{ id: string }>().id;
    const url = new URL(`/api/checkout/create-checkout-session/${id}`, config.baseUrl);

    useEffect(() => {
        fetch(url, {
            method: 'POST',
        })
            .then(res => res.json())
            .then(data => {
                setClientSecret(data.session.client_secret);
            })
            .catch(error => {
                console.error('Error fetching checkout session:', error);
            });
    }, []);

    return (
        <div className={utils.checkout}>
            <h1 className={utils.subHeading}>Checkout Form:</h1>
            <div id="checkout">
                {clientSecret && (
                    <EmbeddedCheckoutProvider stripe={stripePromise} options={{ clientSecret }}>
                        <EmbeddedCheckout />
                    </EmbeddedCheckoutProvider>
                )}
            </div>
        </div>
    );
};

export default CheckoutForm;
