import utils from '../styles/Utils.module.css';

import Error from '../layouts/Error';
import Loading from '../layouts/Loading';
import Contact from '../templates/Contact';
import Displays from '../templates/Displays';
import Carousel from '../templates/Carousel';

import { useGetResource } from '../hooks/genericHooks';
import { Carousel as CarouselType } from '../types/miscTypes';

const Home: React.FC = () => {
    const { data: carousels, isLoading, isError } = useGetResource<CarouselType[]>('carousels');

    const images = carousels?.map(carousel => carousel.image) || [];

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

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
