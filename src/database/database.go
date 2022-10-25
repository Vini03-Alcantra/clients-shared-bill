package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func loadVariablesEnvironment() string {
	var StringConnectionDatabase = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)

	return StringConnectionDatabase
}

func connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var StringConnectionDatabase = loadVariablesEnvironment()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  StringConnectionDatabase,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GetDatabaseConnection() *gorm.DB {
	sqlDB, err := connect()
	if err != nil {
		panic("Failed to create connection with database")
	}

	return sqlDB
}
