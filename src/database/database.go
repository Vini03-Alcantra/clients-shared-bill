package database

import (
	"fmt"
	"log"
	"time"

	"github.com/ClientsSharedBill/src/config"
	"github.com/ClientsSharedBill/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func Connect() error {
	print("Connect String", config.StringDatabaseConnection)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.StringDatabaseConnection,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = db

	return err
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()

	if err != nil {
		return dbConn, err
	}

	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	AutoMigrate()
	return dbConn, nil
}

func AutoMigrate() error {
	// Auto Migrate database
	db, connErr := GetDatabaseConnection()
	if connErr != nil {
		return connErr
	}

	// Add required models here
	err := db.AutoMigrate(&models.Client{})
	fmt.Print("Database Migrated")
	return err
}
