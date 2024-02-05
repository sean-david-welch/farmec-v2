import utils from '../styles/Utils.module.css';
import styles from '../styles/Carousel.module.css';

import CarouselForm from '../forms/CarouselForm';

import { Carousel } from '../types/miscTypes';
import { useGetResource } from '../hooks/genericHooks';
import DeleteButton from './DeleteButton';

interface Props {
    isAdmin: boolean;
}

const CarouselAdmin: React.FC<Props> = ({ isAdmin }) => {
    const { data: carousels } = useGetResource<Carousel[]>('carousels');

    if (!carousels) {
        return (
            <section id="carousel">
                <h1>No models found</h1>
                {isAdmin && <CarouselForm />}
            </section>
        );
    }

    return (
        <section id="carousel">
            <h1 className={utils.sectionHeading}>Carousel:</h1>

            {carousels.map((carousel) => (
                <div key={carousel.id} className={styles.carouselAdmin}>
                    <h1 className={utils.paragraph}>{carousel.name}</h1>
                    {isAdmin && carousel.id && (
                        <div>
                            <CarouselForm id={carousel.id} carousel={carousel} />{' '}
                            <DeleteButton id={carousel.id} resourceKey="carousels" />
                        </div>
                    )}
                </div>
            ))}
            {isAdmin && <CarouselForm />}
        </section>
    );
};

export default CarouselAdmin;
