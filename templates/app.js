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

        statusDiv.onclick = () => {
            showDeviceDetails(device);
        };

        statusDiv.addEventListener('mouseover', () => {
            statusDiv.style.opacity = '0.8';
            statusDiv.style.cursor = 'pointer';
        });

        statusDiv.addEventListener('mouseout', () => {
            statusDiv.style.opacity = '1';
            statusDiv.style.cursor = 'default';
        });

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

function getDeviceRouteKey(device) {
    return device.name || device.hostname || device.deviceid || device.deviceId || '';
}

function hasValidMapPosition(device) {
    if (typeof device.lat !== 'number' || typeof device.lon !== 'number') {
        return false;
    }

    if ((device.lat === 0 && device.lon === 0) || (device.lat === -1 && device.lon === -1)) {
        return false;
    }

    return true;
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
                const routeKey = getDeviceRouteKey(device);
                if (!hasValidMapPosition(device)) {
                    console.log(`Device ${routeKey} has no valid GPS data`);
                    return;
                }

                let marker = L.marker([device.lat, device.lon]);
                marker.bindTooltip(`${routeKey}<br>${device.ip}`, {
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
    const routeKey = getDeviceRouteKey(device);
    if (!routeKey) {
        console.warn('Device has no route key:', device);
        return;
    }

    console.log('Device clicked:', routeKey);
    window.location.href = `/device/${encodeURIComponent(routeKey)}`;
}

document.addEventListener('DOMContentLoaded', function() {
    initializeMap();
    startHealthCheck();
});
