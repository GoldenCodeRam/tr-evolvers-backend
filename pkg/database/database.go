package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file!")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
        os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_LOCAL_PORT"),
	)

    database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

    if err != nil {
        panic("Failed to connect to the database!")
    }

    return database
}
