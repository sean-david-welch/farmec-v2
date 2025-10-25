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
  GoogleMap
};
