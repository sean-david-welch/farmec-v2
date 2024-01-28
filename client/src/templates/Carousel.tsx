import styles from '../styles/Carousel.module.css';
import utils from '../styles/Utils.module.css';

import TypewriterComponent from 'typewriter-effect';

import { useState } from 'react';

interface ImageProps {
    images: string[];
}

const Carousel = (props: ImageProps) => {
    const [currentIndex, setCurrentIndex] = useState(0);

    const nextStep = () => {
        setCurrentIndex(prevIndex => (prevIndex === props.images.length - 1 ? 0 : prevIndex + 1));
    };

    const prevStep = () => {
        setCurrentIndex(prevIndex => (prevIndex === 0 ? props.images.length - 1 : prevIndex - 1));
    };

    return (
        <section id="Hero">
            <div className={styles.heroContainer}>
                <div className={styles.slideshow}>
                    {props.images.map((image, index) => (
                        <img
                            key={index}
                            src={image}
                            alt="Slide"
                            className={`${styles.slides} ${currentIndex === index ? styles.fadeIn : styles.fadeOut}`}
                        />
                    ))}
                    <button className={styles.prevButton} onClick={prevStep} aria-label="last slide">
                        <div>
                            <img src="/icons/chevron-left.svg" alt="Previous Icon" />
                        </div>
                    </button>
                    <button className={styles.nextButton} onClick={nextStep} aria-label="next slide">
                        <div>
                            <img src="/icons/chevron-right.svg" alt="Next Icon" />
                        </div>
                    </button>
                </div>
            </div>
            <div className={utils.typewriter}>
                <h1>
                    <TypewriterComponent
                        options={{
                            loop: false,
                            cursor: '',
                            delay: 50,
                        }}
                        onInit={typewriter => {
                            typewriter
                                .stop()
                                .typeString('Importers & Distributors of Quality Agricultural Machinery')
                                .start();
                        }}
                    />
                </h1>

                <button className={utils.btn}>
                    <a href="#Info">
                        Find Out More: <img src="/icons/chevron-down.svg" alt="down" />
                    </a>
                </button>
            </div>
        </section>
    );
};

export default Carousel;
