package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/goldencoderam/tr-evolvers-backend/pkg/grpc/service"
	"github.com/joho/godotenv"
)

func connect() *gorm.DB {
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

func AutoMigrate() error {
	conn := connect()

	errors := []error{
		conn.AutoMigrate(&EnergyMeterBrand{}),
		conn.AutoMigrate(&EnergyMeter{}),
		conn.AutoMigrate(&EnergyMeterInstallation{}),
	}

	for _, error := range errors {
		if error == nil {
			return error
		}
	}

	return nil
}

func CreateEnergyMeter(
	createEnergyMeterDto *service.CreateEnergyMeterDto,
) (*EnergyMeter, error) {
	conn := connect()

	energyMeter := EnergyMeter{
		Serial:             createEnergyMeterDto.GetSerial(),
		EnergyMeterBrandID: uint(createEnergyMeterDto.GetBrandId()),
	}
	result := conn.Create(&EnergyMeter{})

	return &energyMeter, result.Error
}

func CreateEnergyMeterBrand(
	createEnergyMeterBrand *service.CreateEnergyMeterBrandDto,
) (*EnergyMeterBrand, error) {
	conn := connect()

	energyMeterBrand := EnergyMeterBrand{
		Name: createEnergyMeterBrand.GetBrand(),
	}
	result := conn.Create(&energyMeterBrand)

	return &energyMeterBrand, result.Error
}
