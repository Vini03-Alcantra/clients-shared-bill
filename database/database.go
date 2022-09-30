package database

import (
	"fmt"
	"log"
	"time"

	"github.com/carrot/go-base-api/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var user string
var password string
var db string
var host string
var port string
var ssl string
var timezone string
var dbConn *gorm.DB

func init() {
	user = environment.GetEnvVar("POSTGRES_USER")
	password = environment.GetEnvVar("POSTGRES_PASSWORD")
	db = environment.GetEnvVar("POSTGRES_DB")
	host = environment.GetEnvVar("POSTGRES_HOST")
	port = environment.GetEnvVar("POSTGRES_PORT")
	ssl = environment.GetEnvVar("POSTGRES_SSL")
	timezone = environment.GetEnvVar("POSTGRES_TIMEZONE")
}

func GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, db, port, ssl, timezone)
}

func Connect() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(),
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
