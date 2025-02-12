function initMap() {
    function initMap() {
        try {
            const location = {
                lat: 53.49200990196934,
                lng: -6.5423895598058435
            };

            const map = new google.maps.Map(document.getElementById("map"), {
                zoom: 10,
                center: location,
            });

            new google.maps.Marker({
                position: location,
                map: map,
            });
        } catch (error) {
            console.error("Error initializing map:", error);
            document.getElementById("map").innerHTML = "Error loading map";
        }
    }

    window.initMap = initMap;

    // Add error handler for script loading
    window.gm_authFailure = function() {
        document.getElementById("map").innerHTML = "Error loading Google Maps";
    };
}