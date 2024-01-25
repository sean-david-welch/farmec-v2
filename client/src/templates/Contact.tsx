import utils from '../../styles/Utils.module.css';
import Info from './Info';
import ContactForm from '../../../frontend/src/components/ContactForm';

const Contact: React.FC = () => {
    const mapsKey = import.meta.env.NEXT_PUBLIC_MAPS_KEY;

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
