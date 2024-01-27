import utils from '../../styles/Utils.module.css';

import config from '../lib/env';
import ContactForm from '../forms/ContactForm';

import { Link } from 'react-router-dom';

const Contact: React.FC = () => {
    const mapsKey = config.mapsKey;

    return (
        <section id="contact">
            <h1 className={utils.sectionHeading}>Contact Us:</h1>
            <div className={utils.contactSection}>
                <ContactForm />
                <iframe
                    width="600"
                    height="600"
                    title="Farmec Ireland Ltd Location"
                    className={utils.map}
                    loading="lazy"
                    allowFullScreen
                    src={`https://www.google.com/maps/embed/v1/place?q=Farmec%20Ireland%20ltd&key=${mapsKey}`}
                />
                <div className={utils.info}>
                    <div className={utils.infoSection}>
                        <h1 className={utils.subHeading}>Business Information:</h1>
                        <div className={utils.infoItem}>
                            Opening Hours:
                            <br />
                            <span className={utils.infoItemText}>Monday - Friday: 9am - 5:30pm</span>
                        </div>
                        <div className={utils.infoItem}>
                            Telephone:
                            <br />
                            <span className={utils.infoItemText}>
                                <Link to="tel:01 825 9289">01 825 9289</Link>
                            </span>
                        </div>
                        <div className={utils.infoItem}>
                            International:
                            <br />
                            <span className={utils.infoItemText}>
                                <Link to="tel:+353 1 825 9289">+353 1 825 9289</Link>
                            </span>
                        </div>
                        <div className={utils.infoItem}>
                            Email:
                            <br />
                            <span className={utils.infoItemText}>Info@farmec.ie</span>
                        </div>
                        <div className={utils.infoItem}>
                            Address:
                            <br />
                            <span className={utils.infoItemText}>Clonross, Drumree, Co. Meath, A85PK30</span>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
};

export default Contact;
