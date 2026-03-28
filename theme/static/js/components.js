// ===== Toastify HTMX Integration =====
const TOAST_COLOURS = {
    success: '#16a34a',
    error:   '#b91c1c',
    info:    '#2563eb',
};

document.addEventListener('showToast', function (e) {
    const { message, type = 'success' } = e.detail;
    Toastify({
        text: message,
        duration: 3500,
        gravity: 'top',
        position: 'center',
        stopOnFocus: true,
        style: {
            background: TOAST_COLOURS[type] ?? TOAST_COLOURS.success,
            padding: '16px 28px',
            fontSize: '1.1rem',
            minWidth: '320px',
            borderRadius: '8px',
        },
    }).showToast();
});

// ===== ToTopButton Component =====
document.addEventListener('DOMContentLoaded', function () {
    const toTopBtn = document.getElementById('to-top-btn');

    if (toTopBtn) {
        // Show button by default
        toTopBtn.style.display = 'block';

        // Scroll to top on click
        toTopBtn.addEventListener('click', function () {
            window.scrollTo({
                top: 0,
                behavior: 'smooth'
            });
        });
    }

    // ===== Mobile Menu Toggle =====
    const mobileMenuToggles = document.querySelectorAll('.mobile-menu-toggle');
    mobileMenuToggles.forEach(toggle => {
        toggle.addEventListener('click', function () {
            const menu = this.nextElementSibling;
            if (menu && menu.classList.contains('mobile-menu')) {
                menu.style.display = menu.style.display === 'none' ? 'block' : 'none';
            }
        });
    });
});

// ===== Google Maps Component =====
// Make initializeMap global for callback
window.initializeMap = function () {
    const mapElement = document.getElementById('google-map');
    if (mapElement && typeof google !== 'undefined' && google.maps) {
        const lat = parseFloat(mapElement.dataset.lat);
        const lng = parseFloat(mapElement.dataset.lng);
        const zoom = parseInt(mapElement.dataset.zoom);

        const map = new google.maps.Map(mapElement, {
            zoom: zoom,
            center: {lat: lat, lng: lng},
        });

        new google.maps.Marker({
            position: {lat: lat, lng: lng},
            map: map,
        });
    }
};
