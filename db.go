package main

import (
	"os"
	"log"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func createDBConnection() *gorm.DB {
	// Define a custom Logger to use with gORM
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,       // Enable color
		},
	)
	// Set up a DSN connection string
	dsn := "host=" + os.Getenv("POSTGRES_HOST") +" user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432 sslmode=disable TimeZone=" + os.Getenv("POSTGRES_TIMEZONE")
	//  Open an connection to DB
	dbConn, _ := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: customLogger,
		},
	)

	return dbConn
}

func migrate(dbConn *gorm.DB) {
	dbConn.AutoMigrate(&user{}, &expense{})
}