// Initialisiere Leaflet-Karte
function initializeMap() {
    var map = L.map('map').setView([48.2, 14.1], 10);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap'
    }).addTo(map);

    loadDevices(map);
}

function startHealthCheck() {
    updateHealthStatus();
    setInterval(updateHealthStatus, 10000); // Alle 10 Sekunden
}

/**
 * Fetches and displays the health status from the API.
 */
function updateHealthStatus() {
    fetch("/api/health")
        .then(res => res.json())
        .then(data => {
            console.log('Health Status:', data);
            displayHealthStatus(data);
        })
        .catch(error => {
            console.error('Fehler beim Laden des Health-Status:', error);
        });
}

/**
 * Displays the health status of devices in the healthStatus container.
 * @param {Array} devices - An array of device objects with active status, hostname, and IP address.
 */
function displayHealthStatus(devices) {
    const healthContainer = document.getElementById('healthStatus');
    healthContainer.innerHTML = '';

    devices.forEach(device => {
        const statusColor = device.active ? '#4CAF50' : '#f44336';
        const statusText = device.active ? 'ONLINE' : 'OFFLINE';

        const statusDiv = document.createElement('div');
        statusDiv.style.cssText = `
            padding: 8px;
            margin-bottom: 8px;
            background-color: ${statusColor};
            color: white;
            border-radius: 4px;
            font-size: 12px;
            font-weight: bold;
        `;

        statusDiv.innerHTML = `
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <span>${device.hostname}</span>
                <span>${statusText}</span>
            </div>
            <div style="font-size: 10px; margin-top: 3px; opacity: 0.9;">
                ${device.ip}
            </div>
        `;

        healthContainer.appendChild(statusDiv);
    });
}

/**
 * Loads devices from the API and displays them on the provided map.
 *
 * This function fetches device data from the "/api/devices" endpoint. It checks the response for errors and processes the JSON data.
 * Each device is validated for valid GPS coordinates before being added to the map as a marker.
 * A tooltip is bound to each marker, and a click event is set to show device details using the showDeviceDetails function.
 *
 * @param {Object} map - The map object where the device markers will be added.
 */
function loadDevices(map) {
    fetch("/api/devices")
        .then(res => {
            if (!res.ok) {
                throw new Error('Fehler beim Laden der Geräte');
            }
            return res.json();
        })
        .then(data => {
            console.log('Geladene Geräte:', data);
            data.forEach(device => {
                // Nur Devices mit gültigen GPS-Koordinaten anzeigen (nicht 0,0)
                if (device.lat === 0 && device.lon === 0) {
                    console.log(`Device ${device.name} hat keine gültigen GPS-Daten`);
                    return;
                }

                let marker = L.marker([device.lat, device.lon]);
                // Tooltip beim Hover
                marker.bindTooltip(`${device.name}<br>${device.ip}`, {
                    permanent: false,
                    direction: 'top',
                    offset: [0, -10]
                });

                marker.addTo(map);

                marker.on('click', function() {
                    showDeviceDetails(device);
                });
            });
        })
        .catch(error => {
            console.error('Fehler beim Laden der Geräte:', error);
        });
}

/**
 * Displays device details and redirects to the device's detail page.
 */
function showDeviceDetails(device) {
    console.log('Device clicked:', device.name);

    // Weiterleitung zur Detail-Seite
    window.location.href = `/device/${encodeURIComponent(device.name)}`;
}

// Initialisiere die App beim Laden der Seite
document.addEventListener('DOMContentLoaded', function() {
    initializeMap();
    startHealthCheck();
});

