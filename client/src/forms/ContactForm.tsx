import utils from '../styles/Utils.module.css';
import styles from '../styles/Home.module.css';
import config from '../lib/env';
import { useState } from 'react';

const ContactForm = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [message, setMessage] = useState('');
    const url = new URL('/api/contact', config.baseUrl);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const body = {
            name: name,
            email: email,
            message: message,
        };

        try {
            const repsonse = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body),
            });

            if (repsonse.ok) {
                setName('');
                setEmail('');
                setMessage('');
            }
        } catch (error) {
            console.error('an error occurred', error);
        }
    };

    return (
        <form onSubmit={handleSubmit} className={styles.contactForm}>
            <label htmlFor="name">Name:</label>
            <input
                type="text"
                id="name"
                name="name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
                placeholder="Your name"
            />
            <label htmlFor="email">Email:</label>
            <input
                type="email"
                id="email"
                name="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
                placeholder="Your email"
            />
            <label htmlFor="message">Message:</label>
            <textarea
                id="message"
                name="message"
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                placeholder="Enter your message here..."
                cols={30}
                rows={11}
                required
            />

            <button className={utils.btnForm} type="submit">
                Submit
            </button>
        </form>
    );
};

export default ContactForm;
