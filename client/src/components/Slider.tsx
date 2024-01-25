import { useState } from 'react';
import styles from '../../styles/Carousel.module.css';

interface Props {
    images: string[];
}

const Slider = (props: Props) => {
    const [currentIndex, setCurrentIndex] = useState(0);

    const nextStep = () => {
        setCurrentIndex(prevIndex => (prevIndex === props.images.length - 1 ? 0 : prevIndex + 1));
    };

    const prevStep = () => {
        setCurrentIndex(prevIndex => (prevIndex === 0 ? props.images.length - 1 : prevIndex - 1));
    };

    return (
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
    );
};

export default Slider;
