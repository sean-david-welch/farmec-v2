import { createSignal, For } from 'solid-js';
import styles from '../../styles/Carousel.module.css';

interface Props {
  images: string[];
}

const Slider = (props: Props) => {
  const [currentIndex, setCurrentIndex] = createSignal(0);

  const nextStep = () => {
    setCurrentIndex(prevIndex => (prevIndex === props.images.length - 1 ? 0 : prevIndex + 1));
  };

  const prevStep = () => {
    setCurrentIndex(prevIndex => (prevIndex === 0 ? props.images.length - 1 : prevIndex - 1));
  };

  return (
    <div class={styles.slideshow}>
      <For each={props.images}>
        {(_, index) => (
          <img
            src={props.images[currentIndex()]}
            alt="slides"
            class={`${styles.slides} ${currentIndex() === index() ? styles.fadeIn : styles.fadeOut}`}
          />
        )}
      </For>
      <button class={styles.prevButton} onClick={prevStep} aria-label="last slide">
        <div>
          <img src="/icons/chevron-left.svg" alt="Previous Icon" />
        </div>
      </button>
      <button class={styles.nextButton} onClick={nextStep} aria-label="next slide">
        <div>
          <img src="/icons/chevron-right.svg" alt="Next Icon" />
        </div>
      </button>
    </div>
  );
};

export default Slider;
