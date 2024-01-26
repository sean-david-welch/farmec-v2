import config from '../utils/env';
import Heading from './Heading';
import Slider from '../components/Slider';

import { Carousel } from '../types/miscTypes';

const carousels: Carousel[] = await fetch(`${config.baseUrl}/api/carousels`).then(carousels => carousels.json());

const images = carousels.map(carousels => carousels.image);

const Hero: React.FC = () => {
    return (
        <section id="Hero">
            <Slider images={images} />
            <Heading />
        </section>
    );
};

export default Hero;
