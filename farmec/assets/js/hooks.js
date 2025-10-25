/**
 * Phoenix LiveView Hooks for Farmec application
 * Converted from React component client-side functionality
 */

/**
 * ScrollToTop Hook
 * Handles smooth scrolling to top of page when button is clicked
 * Used by the ToTopButton component
 */
const ScrollToTop = {
  mounted() {
    this.el.addEventListener("click", () => {
      window.scrollTo({
        top: 0,
        behavior: "smooth"
      });
    });
  }
};

/**
 * Carousel Hook
 * Handles image carousel navigation and automatic slideshow
 * Converted from React Carousel component
 */
const Carousel = {
  mounted() {
    this.currentIndex = 0;
    this.images = this.el.querySelectorAll('.slides');
    this.prevButton = this.el.querySelector('.prev-button');
    this.nextButton = this.el.querySelector('.next-button');

    if (this.images.length === 0) return;

    // Set up navigation button listeners
    this.prevButton?.addEventListener('click', () => this.prevSlide());
    this.nextButton?.addEventListener('click', () => this.nextSlide());

    // Optional: Auto-advance every 5 seconds
    this.startAutoPlay();
  },

  destroyed() {
    // Clean up interval when component is destroyed
    if (this.autoPlayInterval) {
      clearInterval(this.autoPlayInterval);
    }
  },

  nextSlide() {
    this.images[this.currentIndex].classList.remove('fade-in');
    this.images[this.currentIndex].classList.add('fade-out');

    this.currentIndex = (this.currentIndex + 1) % this.images.length;

    this.images[this.currentIndex].classList.remove('fade-out');
    this.images[this.currentIndex].classList.add('fade-in');
  },

  prevSlide() {
    this.images[this.currentIndex].classList.remove('fade-in');
    this.images[this.currentIndex].classList.add('fade-out');

    this.currentIndex = (this.currentIndex - 1 + this.images.length) % this.images.length;

    this.images[this.currentIndex].classList.remove('fade-out');
    this.images[this.currentIndex].classList.add('fade-in');
  },

  startAutoPlay() {
    // Auto-advance every 5 seconds
    this.autoPlayInterval = setInterval(() => {
      this.nextSlide();
    }, 5000);
  }
};

/**
 * GoogleMap Hook
 * Initializes Google Maps for the Map component
 * Requires Google Maps API to be loaded
 */
const GoogleMap = {
  mounted() {
    const lat = parseFloat(this.el.dataset.lat) || 53.49200990196934;
    const lng = parseFloat(this.el.dataset.lng) || -6.5423895598058435;

    const location = { lat, lng };

    // Wait for Google Maps API to be loaded
    if (window.google && window.google.maps) {
      this.initMap(location);
    } else {
      // Retry after a short delay if API not loaded yet
      setTimeout(() => {
        if (window.google && window.google.maps) {
          this.initMap(location);
        }
      }, 100);
    }
  },

  initMap(location) {
    const map = new google.maps.Map(this.el, {
      center: location,
      zoom: 10
    });

    new google.maps.Marker({
      position: location,
      map: map
    });
  }
};

export default {
  ScrollToTop,
  Carousel,
  GoogleMap
};
