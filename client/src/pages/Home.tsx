import utils from '../styles/Utils.module.css';

import ErrorPage from '../layouts/Error';
import Loading from '../layouts/Loading';
import Contact from '../templates/Contact';
import Displays from '../templates/Displays';
import Carousel from '../templates/Carousel';

import {useGetResource} from '../hooks/genericHooks';
import {Carousel as CarouselType} from '../types/miscTypes';
import {Helmet} from "react-helmet";
import {FC} from "react";

const Home: FC = () => {
    const { data: carousels, isLoading, isError } = useGetResource<CarouselType[]>('carousels');

    const images = carousels?.map(carousel => carousel.image) || [];

    if (isError) return <ErrorPage />;
    if (isLoading) return <Loading />;

    return (
        <>
            <Helmet>
                <title>Home - Farmec Ireland</title>
                <meta name="description" content="Welcome to Farmec Ireland Ltd. Importers and Distributors of High quality Farm and Amenity Machinery"/>

                <meta property="og:title" content="Home - Farmec Ireland"/>
                <meta property="og:description" content="Discover Farmec's staff, history, and vision for the future."/>
                <meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"/>
                <meta property="og:url" content="https://www.farmec.ie"/>
                <meta property="og:type" content="website"/>

                <meta name="twitter:card" content="summary_large_image"/>
                <meta name="twitter:title" content="Home - Farmec Ireland"/>
                <meta name="twitter:description" content="Welcome to Farmec Ireland Ltd. Importers and Distributors of High quality Farm and Amenity Machinery"/>
                <meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"/>
                <link rel="canonical" href="https://www.farmec.ie"/>
            </Helmet>
            <section id="Home">
                <div className={utils.home}>
                    <Carousel images={images}/>
                    <Displays/>
                    <Contact/>
                </div>
            </section>
        </>
    );
};

export default Home;
