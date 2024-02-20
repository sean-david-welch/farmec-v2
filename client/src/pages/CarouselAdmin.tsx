import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';

import ErrorPage from '../layouts/Error';
import Loading from '../layouts/Loading';
import CarouselForm from '../forms/CarouselForm';
import DeleteButton from '../components/DeleteButton';

import { Carousel } from '../types/miscTypes';
import { useGetResource } from '../hooks/genericHooks';
import { useUserStore } from '../lib/store';
import { Fragment } from 'react';

const CarouselAdmin: React.FC = () => {
    const { isAdmin } = useUserStore();
    const { data: carousels, isError, isLoading } = useGetResource<Carousel[]>('carousels');

    if (isError) return <ErrorPage />;
    if (isLoading) return <Loading />;

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
            <h1 className={utils.sectionHeading}>Carousels:</h1>
            {isAdmin ? (
                <Fragment>
                    {carousels.map(carousel => (
                        <div key={carousel.id} className={styles.carouselAdmin}>
                            <h1 className={utils.mainHeading}>{carousel.name}</h1>
                            <img src={carousel.image} alt="carousel image" width={400} />
                            {isAdmin && carousel.id && (
                                <div className={utils.optionsBtn}>
                                    <CarouselForm id={carousel.id} carousel={carousel} />{' '}
                                    <DeleteButton id={carousel.id} resourceKey="carousels" />
                                </div>
                            )}
                        </div>
                    ))}
                    <CarouselForm />
                </Fragment>
            ) : null}
        </section>
    );
};

export default CarouselAdmin;
