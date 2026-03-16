
DROP TABLE IF EXISTS dataset CASCADE;
DROP TABLE IF EXISTS device_hardware CASCADE;
DROP TABLE IF EXISTS device CASCADE;
DROP TABLE IF EXISTS cpu CASCADE;
DROP TABLE IF EXISTS ram CASCADE;

CREATE TABLE IF NOT EXISTS cpu (
                                   cpuID SERIAL PRIMARY KEY,
                                   cores INT NOT NULL,
                                   threads INT NOT NULL,
                                   model VARCHAR(255) NOT NULL,
    speed BIGINT
    );

CREATE TABLE IF NOT EXISTS ram (
                                   ramID SERIAL PRIMARY KEY,
                                   capacity BIGINT NOT NULL,
                                   model VARCHAR(255),
    speed BIGINT
    );

CREATE TABLE IF NOT EXISTS device (
                                      deviceID VARCHAR(255) PRIMARY KEY,
    os VARCHAR(255),
    hostname VARCHAR(255),
    ip_address VARCHAR(255),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS device_hardware (
                                               deviceID VARCHAR(255) PRIMARY KEY,
    cpuID INT NOT NULL,
    ramID INT NOT NULL,
    FOREIGN KEY (deviceID) REFERENCES device(deviceID) ON DELETE CASCADE,
    FOREIGN KEY (cpuID) REFERENCES cpu(cpuID),
    FOREIGN KEY (ramID) REFERENCES ram(ramID)
    );

CREATE TABLE IF NOT EXISTS dataset (
                                       datasetID SERIAL PRIMARY KEY,
                                       timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                       deviceID VARCHAR(255) NOT NULL,

    -- CPU Messwerte
    cpuUsage DECIMAL(5, 2),
    cpuTemp DECIMAL(5, 2),

    -- RAM Messwerte
    ramUsed BIGINT,
    ramAvailable BIGINT,
    ramUsedPercent DECIMAL(5, 2),

    -- GPS Daten (ändern sich bei jedem Update)
    gpsLatitude DECIMAL(10, 8),
    gpsLongitude DECIMAL(11, 8),
    gpsAltitude DECIMAL(10, 2),
    gpsAccuracy DECIMAL(10, 2),
    gpsCity VARCHAR(255),
    gpsCountry VARCHAR(255),
    gpsRegion VARCHAR(255),

    FOREIGN KEY (deviceID) REFERENCES device(deviceID) ON DELETE CASCADE
    );

-- Indizes für Performance
CREATE INDEX idx_dataset_deviceID ON dataset(deviceID);
CREATE INDEX idx_dataset_timestamp ON dataset(timestamp);
CREATE INDEX idx_device_hardware_cpu ON device_hardware(cpuID);
CREATE INDEX idx_device_hardware_ram ON device_hardware(ramID);

