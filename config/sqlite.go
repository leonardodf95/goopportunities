package config

import (
	"os"

	"github.com/leonardodf95/goopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("SQLite database does not exist, creating it")

		//CREATE DATABASE FILE
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating SQLite database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("Error creating SQLite database: %v", err)
			return nil, err
		}

		file.Close()
	}

	//CREATE DATABASE AND CONNECT
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing SQLite database: %v", err)
		return nil, err
	}

	//MIGRATE SCHEMA
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err)
		return nil, err
	}

	logger.Info("SQLite database initialized")
	return db, nil
}
