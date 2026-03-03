package config

import (
	"fmt"
	"os"
	"os/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")

	var db *gorm.DB
	var err error

	if dbURL != "" {
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&user.User{}); err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully")

	return db, nil
}
