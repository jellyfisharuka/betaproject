package db

import (
	"betaproject/internal/config"
	"log"
	//"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDB() {
	var err error
	//dsn := os.Getenv("DB")
	dsn := config.GetEnvConfig().DbDSN
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	log.Println("Successfully connected to database")

	err = DB.AutoMigrate()
	if err != nil {
		panic("Failed to migrate DB schemas")
	}

}