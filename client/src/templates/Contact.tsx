import utils from '../../styles/Utils.module.css';

import config from '../lib/env';

import Info from './Info';
import ContactForm from '../forms/ContactForm';

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
                <Info />
            </div>
        </section>
    );
};

export default Contact;
