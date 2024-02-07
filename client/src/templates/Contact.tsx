import utils from '../styles/Utils.module.css';
import styles from '../styles/Home.module.css';

import config from '../lib/env';
import ContactForm from '../forms/ContactForm';

import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFacebook } from '@fortawesome/free-brands-svg-icons/faFacebook';
import { faTwitter } from '@fortawesome/free-brands-svg-icons/faTwitter';

const Contact: React.FC = () => {
    const mapsKey = config.mapsKey;

    return (
        <section id="contact">
            <h1 className={utils.sectionHeading}>Contact Us:</h1>
            <div className={styles.contactSection}>
                <ContactForm />
                <iframe
                    width="600"
                    height="600"
                    title="Farmec Ireland Ltd Location"
                    className={styles.map}
                    loading="lazy"
                    allowFullScreen
                    src={`https://www.google.com/maps/embed/v1/place?q=Farmec%20Ireland%20ltd&key=${mapsKey}`}
                />

                <div className={styles.infoSection}>
                    <h1 className={styles.subHeading}>Business Information:</h1>
                    <div className={styles.info}>
                        <div className={styles.infoItem}>
                            Opening Hours:
                            <br />
                            <span className={styles.infoItemText}>Monday - Friday: 9am - 5:30pm</span>
                        </div>
                        <div className={styles.infoItem}>
                            Telephone:
                            <br />
                            <span className={styles.infoItemText}>
                                <Link to="tel:01 825 9289">01 825 9289</Link>
                            </span>
                        </div>
                        <div className={styles.infoItem}>
                            International:
                            <br />
                            <span className={styles.infoItemText}>
                                <Link to="tel:+353 1 825 9289">+353 1 825 9289</Link>
                            </span>
                        </div>
                        <div className={styles.infoItem}>
                            Email:
                            <br />
                            <span className={styles.infoItemText}>Info@farmec.ie</span>
                        </div>
                        <div className={styles.infoItem}>
                            Address:
                            <br />
                            <span className={styles.infoItemText}>Clonross, Drumree, Co. Meath, A85PK30</span>
                        </div>

                        <div className={styles.infoItem}>
                            <div className={styles.socialLinks}>
                                <Link
                                    className={styles.socials}
                                    to={'https://www.facebook.com/FarmecIreland/'}
                                    target={'_blank'}
                                    rel={'noopener noreferrer'}
                                    aria-label="Visit our Facebook page"
                                >
                                    <FontAwesomeIcon icon={faFacebook} />
                                </Link>
                                <Link
                                    className={styles.socials}
                                    to={'https://twitter.com/farmec1?lang=en'}
                                    target={'_blank'}
                                    rel={'noopener noreferrer'}
                                    aria-label="Visit our Twiiter page"
                                >
                                    <FontAwesomeIcon icon={faTwitter} />
                                </Link>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
};

export default Contact;
