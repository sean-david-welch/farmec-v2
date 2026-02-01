// Get CSRF token from DOM
function getCookie(name) {
    let cookieValue = null;
    if (document.cookie && document.cookie !== '') {
        const cookies = document.cookie.split(';');
        for (let i = 0; i < cookies.length; i++) {
            const cookie = cookies[i].trim();
            if (cookie.substring(0, name.length + 1) === (name + '=')) {
                cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
                break;
            }
        }
    }
    return cookieValue;
}

const csrftoken = getCookie('csrftoken');

// ===== ToTopButton Component =====
document.addEventListener('DOMContentLoaded', function() {
    const toTopBtn = document.getElementById('to-top-btn');

    if (toTopBtn) {
        // Show/hide button on scroll
        window.addEventListener('scroll', function() {
            if (window.pageYOffset > 300) {
                toTopBtn.style.display = 'block';
            } else {
                toTopBtn.style.display = 'none';
            }
        });

        // Scroll to top on click
        toTopBtn.addEventListener('click', function() {
            window.scrollTo({
                top: 0,
                behavior: 'smooth'
            });
        });
    }

    // ===== DeleteButton Component =====
    const deleteButtons = document.querySelectorAll('.delete-btn');
    deleteButtons.forEach(btn => {
        btn.addEventListener('click', async function(e) {
            e.preventDefault();

            const resourceType = this.dataset.resourceType;
            const resourceId = this.dataset.resourceId;
            const navigateBack = this.dataset.navigateBack === 'true';

            // Show confirmation
            if (!confirm(`Are you sure you want to delete this ${resourceType}?`)) {
                return;
            }

            try {
                // Make DELETE request to Go API
                const response = await fetch(`/api/${resourceType}/${resourceId}`, {
                    method: 'DELETE',
                    headers: {
                        'X-CSRFToken': csrftoken,
                        'Content-Type': 'application/json',
                    }
                });

                if (response.ok) {
                    alert(`${resourceType} deleted successfully`);
                    if (navigateBack) {
                        window.location.href = '/';
                    } else {
                        location.reload();
                    }
                } else {
                    const errorData = await response.json();
                    alert(`Error: ${errorData.message || 'Failed to delete'}`);
                    console.error('Delete error:', errorData);
                }
            } catch (error) {
                console.error('Error deleting resource:', error);
                alert('An error occurred while deleting');
            }
        });
    });

    // ===== DownloadPdf Component =====
    const pdfButtons = document.querySelectorAll('.pdf-download-btn');
    pdfButtons.forEach(btn => {
        btn.addEventListener('click', async function(e) {
            e.preventDefault();

            const pdfType = this.dataset.pdfType;
            const warrantyClaimId = this.dataset.warrantyClaimId;
            const registrationId = this.dataset.registrationId;

            try {
                // Build request body
                const body = {
                    pdf_type: pdfType,
                };

                if (warrantyClaimId) {
                    body.warranty_claim_id = warrantyClaimId;
                }
                if (registrationId) {
                    body.registration_id = registrationId;
                }

                // Make POST request to Go API
                const response = await fetch(`/api/pdf/${pdfType}`, {
                    method: 'POST',
                    headers: {
                        'X-CSRFToken': csrftoken,
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(body)
                });

                if (response.ok) {
                    // Get the blob and download
                    const blob = await response.blob();
                    const url = window.URL.createObjectURL(blob);
                    const a = document.createElement('a');
                    a.href = url;
                    a.download = `${pdfType}_document.pdf`;
                    document.body.appendChild(a);
                    a.click();
                    window.URL.revokeObjectURL(url);
                    document.body.removeChild(a);
                } else {
                    const errorData = await response.json();
                    alert(`Error: ${errorData.message || 'Failed to generate PDF'}`);
                    console.error('PDF generation error:', errorData);
                }
            } catch (error) {
                console.error('Error downloading PDF:', error);
                alert('An error occurred while generating PDF');
            }
        });
    });

    // ===== Mobile Menu Toggle =====
    const mobileMenuToggles = document.querySelectorAll('.mobile-menu-toggle');
    mobileMenuToggles.forEach(toggle => {
        toggle.addEventListener('click', function() {
            const menu = this.nextElementSibling;
            if (menu && menu.classList.contains('mobile-menu')) {
                menu.style.display = menu.style.display === 'none' ? 'block' : 'none';
            }
        });
    });
});

// ===== Google Maps Component =====
if (typeof google !== 'undefined' && google.maps) {
    document.addEventListener('DOMContentLoaded', function() {
        const mapElement = document.getElementById('google-map');
        if (mapElement) {
            const lat = parseFloat(mapElement.dataset.lat);
            const lng = parseFloat(mapElement.dataset.lng);
            const zoom = parseInt(mapElement.dataset.zoom);

            const map = new google.maps.Map(mapElement, {
                zoom: zoom,
                center: { lat: lat, lng: lng },
            });

            new google.maps.Marker({
                position: { lat: lat, lng: lng },
                map: map,
            });
        }
    });
}
