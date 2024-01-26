import Carousel from '../types/carousel';

import Heading from './Heading';
import Slider from '../components/Slider';
import config from '../utils/env';

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
