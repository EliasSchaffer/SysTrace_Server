-- CPU Tabelle
CREATE TABLE IF NOT EXISTS cpu (
    cpuID SERIAL PRIMARY KEY,
    cores INT NOT NULL,
    threads INT NOT NULL,
    model VARCHAR(255) NOT NULL,
    usage DECIMAL(5, 2),
    temp DECIMAL(5, 2)
);

-- RAM Tabelle
CREATE TABLE IF NOT EXISTS ram (
    ramID SERIAL PRIMARY KEY,
    total BIGINT NOT NULL,
    used BIGINT NOT NULL,
    available BIGINT NOT NULL,
    usedPercent DECIMAL(5, 2),
    model VARCHAR(255),
    speed BIGINT
);

-- GPS Tabelle
CREATE TABLE IF NOT EXISTS gps (
    gpsID SERIAL PRIMARY KEY,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    altitude DECIMAL(10, 2),
    accuracy DECIMAL(10, 2),
    city VARCHAR(255),
    country VARCHAR(255),
    region VARCHAR(255)
);

-- Device Tabelle
CREATE TABLE IF NOT EXISTS device (
    deviceID VARCHAR(255) PRIMARY KEY,
    ramID INT NOT NULL,
    cpuID INT NOT NULL,
    gpsID INT NOT NULL,
    os VARCHAR(255),
    hostname VARCHAR(255),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ramID) REFERENCES ram(ramID),
    FOREIGN KEY (cpuID) REFERENCES cpu(cpuID),
    FOREIGN KEY (gpsID) REFERENCES gps(gpsID)
);

-- Dataset Tabelle
CREATE TABLE IF NOT EXISTS dataset (
    datasetID SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deviceID VARCHAR(255) NOT NULL,
    FOREIGN KEY (deviceID) REFERENCES device(deviceID)
);

