import utils from '../styles/Utils.module.css';

const ToTopButton: React.FC = () => {
    const scrollToTop = () => {
        window.scrollTo({
            top: 0,
            behavior: 'smooth',
        });
    };

    return (
        <button id="toTopButton" aria-label="scroll-to-top-button" className={utils.toTopButton} onClick={scrollToTop}>
            <img src="/icons/arrow-up.svg" alt="arrow-up" />
        </button>
    );
};

export default ToTopButton;
