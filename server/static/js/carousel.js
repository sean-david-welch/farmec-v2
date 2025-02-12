// public/js/carousel.js
class Carousel {
    constructor() {
        this.currentIndex = 0;
        this.slides = document.querySelectorAll('.slides');
        this.prevButton = document.getElementById('prevSlide');
        this.nextButton = document.getElementById('nextSlide');
        this.typewriterText = document.getElementById('typewriterText');

        this.init();
    }

    init() {
        // Set initial state
        this.updateSlides();
        this.setupEventListeners();
        this.initTypewriter();
    }

    updateSlides() {
        this.slides.forEach((slide, index) => {
            if (index === this.currentIndex) {
                slide.classList.add('fadeIn');
                slide.classList.remove('fadeOut');
            } else {
                slide.classList.add('fadeOut');
                slide.classList.remove('fadeIn');
            }
        });
    }

    nextStep() {
        this.currentIndex = this.currentIndex === this.slides.length - 1 ? 0 : this.currentIndex + 1;
        this.updateSlides();
    }

    prevStep() {
        this.currentIndex = this.currentIndex === 0 ? this.slides.length - 1 : this.currentIndex - 1;
        this.updateSlides();
    }

    setupEventListeners() {
        this.prevButton.addEventListener('click', () => this.prevStep());
        this.nextButton.addEventListener('click', () => this.nextStep());

        // Handle image errors
        this.slides.forEach(slide => {
            slide.addEventListener('error', (e) => {
                e.target.src = '/default.jpg';
            });
        });
    }

    initTypewriter() {
        const text = 'Importers & Distributors of Quality Agricultural Machinery';
        let index = 0;
        const speed = 50;

        const type = () => {
            if (index < text.length) {
                this.typewriterText.textContent += text.charAt(index);
                index++;
                setTimeout(type, speed);
            }
        };

        type();
    }
}

// Initialize carousel when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new Carousel();
});