package database

import (
	"SysTrace_Server/data/static"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "systrace"
	password = "systrace_secure_password"
	dbname   = "systrace_db"
)

func InitDatabase() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("Error openng connection to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("Coudnt connect to Database: %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	log.Println("Database Connected!")
	return nil
}

func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

func IsConnected() bool {
	if DB == nil {
		return false
	}
	return DB.Ping() == nil
}

func InsertFullDataSet(hostname string, device static.Device) error {

	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("Error starting transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var cpuID int
	err = tx.QueryRow(
		"SELECT cpuID FROM cpu WHERE cores = $1 AND threads = $2 AND model = $3",
		device.Hardware.CPU.Cores, device.Hardware.CPU.Threads, device.Hardware.CPU.Model,
	).Scan(&cpuID)

	if err == sql.ErrNoRows {
		err = tx.QueryRow(
			"INSERT INTO cpu (cores, threads, model) VALUES ($1, $2, $3) RETURNING cpuID",
			device.Hardware.CPU.Cores, device.Hardware.CPU.Threads, device.Hardware.CPU.Model,
		).Scan(&cpuID)
		if err != nil {
			return fmt.Errorf("Error inserting CPU: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("Error checking CPU: %v", err)
	}

	var ramID int
	err = tx.QueryRow(
		"SELECT ramID FROM ram WHERE capacity = $1 AND model = $2",
		device.Hardware.MEMORY.Total, device.Hardware.MEMORY.Model,
	).Scan(&ramID)

	if err == sql.ErrNoRows {
		err = tx.QueryRow(
			"INSERT INTO ram (capacity, model, speed) VALUES ($1, $2, $3) RETURNING ramID",
			device.Hardware.MEMORY.Total, device.Hardware.MEMORY.Model, device.Hardware.MEMORY.Speed,
		).Scan(&ramID)
		if err != nil {
			return fmt.Errorf("Error inserting RAM: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("Error checking RAM: %v", err)
	}

	var existingDeviceID string
	err = tx.QueryRow("SELECT deviceID FROM device WHERE deviceID = $1", device.ID).Scan(&existingDeviceID)

	if err == sql.ErrNoRows {
		_, err = tx.Exec(
			"INSERT INTO device (deviceID, os, hostname, createdAt) VALUES ($1, $2, $3, $4)",
			device.ID, device.OS, device.Hostname, time.Now(),
		)
		if err != nil {
			return fmt.Errorf("Error inserting Device: %v", err)
		}

		// Erstelle device_hardware Beziehung
		_, err = tx.Exec(
			"INSERT INTO device_hardware (deviceID, cpuID, ramID) VALUES ($1, $2, $3)",
			device.ID, cpuID, ramID,
		)
		if err != nil {
			return fmt.Errorf("Error inserting Device Hardware: %v", err)
		}
		fmt.Printf("Neues Device %s erstellt\n", device.ID)
	} else if err != nil {
		return fmt.Errorf("Error checking device: %v", err)
	}

	// 4. Erstelle Dataset mit Messwerten
	_, err = tx.Exec(`
		INSERT INTO dataset (
			timestamp, deviceID,
			cpuUsage, cpuTemp,
			ramUsed, ramAvailable, ramUsedPercent,
			gpsLatitude, gpsLongitude, gpsAltitude, gpsAccuracy, gpsCity, gpsCountry, gpsRegion
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		time.Now(), device.ID,
		device.Hardware.CPU.Usage, device.Hardware.CPU.Temp,
		device.Hardware.MEMORY.Used, device.Hardware.MEMORY.Available, device.Hardware.MEMORY.UsedPercent,
		device.GPS.Latitude, device.GPS.Longitude, device.GPS.Altitude, device.GPS.Accuracy, device.GPS.City, device.GPS.Country, device.GPS.Region,
	)
	if err != nil {
		return fmt.Errorf("Error inserting Dataset: %v", err)
	}
	fmt.Printf("Dataset für Device %s erstellt\n", device.ID)

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("Error committing transaction: %v", err)
	}

	return nil

}
