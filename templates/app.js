// Initialisiere Leaflet-Karte
function initializeMap() {
    var map = L.map('map').setView([48.2, 14.1], 10);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap'
    }).addTo(map);

    loadDevices(map);
}

function loadDevices(map) {
    fetch("/devices")
        .then(res => {
            if (!res.ok) {
                throw new Error('Fehler beim Laden der Geräte');
            }
            return res.json();
        })
        .then(data => {
            console.log('Geladene Geräte:', data);
            data.forEach(device => {
                L.marker([device.lat, device.lon])
                    .addTo(map)
                    .bindPopup(`<strong>${device.name}</strong>`);
            });
        })
        .catch(error => {
            console.error('Fehler beim Laden der Geräte:', error);
        });
}

// Initialisiere die App beim Laden der Seite
document.addEventListener('DOMContentLoaded', function() {
    initializeMap();
});

