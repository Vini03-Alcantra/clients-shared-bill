package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ClientsSharedBill/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB
var StringConexaoBanco = ""

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	StringConexaoBanco = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  StringConexaoBanco,
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
