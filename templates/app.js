// Initialize Leaflet map
function initializeMap() {
    var map = L.map('map').setView([48.2, 14.1], 10);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap'
    }).addTo(map);

    loadDevices(map);
}

function startHealthCheck() {
    updateHealthStatus();
    setInterval(updateHealthStatus, 10000);
}

function updateHealthStatus() {
    fetch("/api/health")
        .then(res => res.json())
        .then(data => {
            console.log('Health Status:', data);
            displayHealthStatus(data);
        })
        .catch(error => {
            console.error('Error loading health status:', error);
        });
}

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

function loadDevices(map) {
    fetch("/api/devices")
        .then(res => {
            if (!res.ok) {
                throw new Error('Error loading devices');
            }
            return res.json();
        })
        .then(data => {
            console.log('Loaded devices:', data);
            data.forEach(device => {
                if (device.lat === 0 && device.lon === 0) {
                    console.log(`Device ${device.name} has no valid GPS data`);
                    return;
                }

                let marker = L.marker([device.lat, device.lon]);
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
            console.error('Error loading devices:', error);
        });
}

function showDeviceDetails(device) {
    console.log('Device clicked:', device.name);

    window.location.href = `/device/${encodeURIComponent(device.name)}`;
}

document.addEventListener('DOMContentLoaded', function() {
    initializeMap();
    startHealthCheck();
});
