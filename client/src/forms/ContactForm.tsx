import utils from '../styles/Utils.module.css';
import styles from '../styles/Home.module.css';
import config from '../lib/env';

const ContactForm = () => {
    const url = new URL('/api/contact', config.baseUrl);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);

        const body = {
            name: formData.get('name'),
            email: formData.get('email'),
            message: formData.get('message'),
        };

        // Add error handling for the fetch request as needed
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        });

        // Handle response (e.g., showing a success message)
    };

    return (
        <form onSubmit={handleSubmit} className={styles.contactForm}>
            <label htmlFor="name">Name:</label>
            <input type="text" name="name" required={true} placeholder="name" />
            <label htmlFor="email">Email:</label>
            <input type="email" name="email" required={true} placeholder="email" />
            <label htmlFor="message">Message:</label>
            <textarea
                name="message"
                placeholder="Enter your message here..."
                cols={30}
                rows={11}
                required={true}
            />

            <div className={styles.recaptcha}></div>

            <button className={utils.btnForm} type="submit">
                Submit
            </button>
        </form>
    );
};

export default ContactForm;
