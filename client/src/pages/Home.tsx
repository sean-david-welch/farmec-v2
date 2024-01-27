import utils from '../styles/Utils.module.css';

import Contact from '../templates/Contact';
import Displays from '../templates/Displays';
import Carousel from '../templates/Carousel';

import { useGetResource } from '../hooks/genericHooks';
import { Carousel as CarouselType } from '../types/miscTypes';

const Home: React.FC = () => {
    const carousels = useGetResource<CarouselType[]>('carousels');

    const images = carousels.data?.map(carousel => carousel.image) || [];

    return (
        <section id="Home">
            <div className={utils.home}>
                <Carousel images={images} />
                <Displays />
                <Contact />
            </div>
        </section>
    );
};

export default Home;
