// ===== Toastify Integration =====
const TOAST_COLOURS = {
    success: '#16a34a',
    error:   '#b91c1c',
    info:    '#2563eb',
};

window.showToast = function (message, type = 'success') {
    Toastify({
        text: message,
        duration: 6000,
        gravity: 'top',
        position: 'center',
        stopOnFocus: true,
        close: true,
        style: {
            background: TOAST_COLOURS[type] ?? TOAST_COLOURS.success,
            padding: '16px 28px',
            fontSize: '1.1rem',
            minWidth: '320px',
            borderRadius: '8px',
        },
    }).showToast();
};

document.addEventListener('showToast', function (e) {
    const { message, type = 'success' } = e.detail;
    window.showToast(message, type);
});

// ===== Scroll Reveal =====
const scrollRevealObserver = new IntersectionObserver((entries) => {
    entries.forEach((entry, i) => {
        if (entry.isIntersecting) {
            setTimeout(() => entry.target.classList.add('visible'), i * 100);
            scrollRevealObserver.unobserve(entry.target);
        }
    });
}, { threshold: 0.15 });

document.querySelectorAll('.scroll-reveal').forEach(el => scrollRevealObserver.observe(el));

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

});
