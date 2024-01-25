import type Carousel from '../../../frontend/src/types/carousel';

import Slider from '../../../frontend/src/components/Slider';
import Heading from './Heading';

const carousels: Carousel[] = await fetch('http://localhost:4321/api/carousels').then(carousels => carousels.json());

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
