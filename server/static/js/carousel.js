document.addEventListener('DOMContentLoaded', () => {
    // State
    let currentIndex = 0;
    const slides = document.querySelectorAll('.slides');
    const prevButton = document.getElementById('prevSlide');
    const nextButton = document.getElementById('nextSlide');
    const typewriterText = document.getElementById('typewriterText');

    // Functions
    function updateSlides() {
        slides.forEach((slide, index) => {
            if (index === currentIndex) {
                slide.classList.add('fadeIn');
                slide.classList.remove('fadeOut');
            } else {
                slide.classList.add('fadeOut');
                slide.classList.remove('fadeIn');
            }
        });
    }

    function nextStep() {
        currentIndex = currentIndex === slides.length - 1 ? 0 : currentIndex + 1;
        updateSlides();
    }

    function prevStep() {
        currentIndex = currentIndex === 0 ? slides.length - 1 : currentIndex - 1;
        updateSlides();
    }

    function setupEventListeners() {
        prevButton.addEventListener('click', prevStep);
        nextButton.addEventListener('click', nextStep);

        // Handle image errors
        slides.forEach(slide => {
            slide.addEventListener('error', (e) => {
                e.target.src = '/default.jpg';
            });
        });
    }

    function initTypewriter() {
        const text = 'Importers & Distributors of Quality Agricultural Machinery';
        let index = 0;
        const speed = 50;

        function type() {
            if (index < text.length) {
                typewriterText.textContent += text.charAt(index);
                index++;
                setTimeout(type, speed);
            }
        }

        type();
    }

    // Initialize
    updateSlides();
    setupEventListeners();
    initTypewriter();
});