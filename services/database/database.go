package database

import (
	"database/sql"
	"fmt"
	"log"

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
