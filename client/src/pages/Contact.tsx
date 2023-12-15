import { Component } from 'solid-js';
import utils from '../../styles/Utils.module.css';
import Info from './Info';
import ContactForm from '../components/ContactForm';

const Contact: Component = () => {
  const mapsKey = import.meta.env.NEXT_PUBLIC_MAPS_KEY;

  return (
    <section id="contact">
      <h1 class={utils.sectionHeading}>Contact Us:</h1>
      <div class={utils.contactSection}>
        <ContactForm />
        <iframe
          width="600"
          height="600"
          title="Farmec Ireland Ltd Location"
          class={utils.map}
          loading="lazy"
          allowfullscreen
          src={`https://www.google.com/maps/embed/v1/place?q=Farmec%20Ireland%20ltd&key=${mapsKey}`}
        />
        <Info />
      </div>
    </section>
  );
};

export default Contact;
