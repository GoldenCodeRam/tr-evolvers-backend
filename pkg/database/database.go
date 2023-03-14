package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/goldencoderam/tr-evolvers-backend/internal/grpc/service"
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
	createEnergyMeterDto *service.CreateDto,
) (*EnergyMeter, error) {
	conn := connect()

	energyMeter := EnergyMeter{
		Serial:             createEnergyMeterDto.GetSerial(),
		EnergyMeterBrandID: uint(createEnergyMeterDto.GetBrandId()),
	}
	result := conn.Create(&EnergyMeter{})

	return &energyMeter, result.Error
}

func GetEnergyMeter(
	energyMeterDto *service.Dto,
) (*EnergyMeter, error) {
	conn := connect()

	var energyMeter EnergyMeter
	result := conn.First(&energyMeter, energyMeterDto.GetSerial())
	return &energyMeter, result.Error
}

func DeleteEnergyMeter(
	energyMeterDto *service.Dto,
) (*EnergyMeter, error) {
	conn := connect()

	var energyMeter EnergyMeter
	result := conn.Delete(&energyMeter, energyMeterDto.GetSerial())
	return &energyMeter, result.Error
}

func CreateEnergyMeterBrand(
	createEnergyMeterBrandDto *service.CreateBrandDto,
) (*EnergyMeterBrand, error) {
	conn := connect()

	energyMeterBrand := EnergyMeterBrand{
		Name: createEnergyMeterBrandDto.GetBrand(),
	}
	result := conn.Create(&energyMeterBrand)

	return &energyMeterBrand, result.Error
}

func GetEnergyMeterBrand(
	energyMeterBrandDto *service.BrandDto,
) (*EnergyMeterBrand, error) {
	conn := connect()

	var energyMeterBrand EnergyMeterBrand
	result := conn.First(&energyMeterBrand, energyMeterBrandDto.GetId())
	return &energyMeterBrand, result.Error
}

func DeleteEnergyMeterBrand(
	energyMeterBrandDto *service.BrandDto,
) (*EnergyMeterBrand, error) {
	conn := connect()

	var energyMeterBrand EnergyMeterBrand
	result := conn.Delete(&energyMeterBrand, energyMeterBrandDto.GetId())
	return &energyMeterBrand, result.Error
}

func CreateEnergyMeterInstallation(
	createEnergyMeterInstallationDto *service.CreateInstallationDto,
) (*EnergyMeterInstallation, error) {
	conn := connect()

	installationDate := createEnergyMeterInstallationDto.GetInstallationDate().AsTime()
	retirementDate := createEnergyMeterInstallationDto.GetRetirementDate().AsTime()
	energyMeterInstallation := EnergyMeterInstallation{
		Address:          createEnergyMeterInstallationDto.GetAddress(),
		InstallationDate: &installationDate,
		RetirementDate:   &retirementDate,
		Lines:            uint8(createEnergyMeterInstallationDto.GetLines()),
		IsActive:         createEnergyMeterInstallationDto.GetIsActive(),
		CreatedAt:        &time.Time{},

		EnergyMeterSerial: createEnergyMeterInstallationDto.GetSerialId(),
	}
	result := conn.Create(&energyMeterInstallation)

	return &energyMeterInstallation, result.Error
}

func GetEnergyMeterInstallation(
	energyMeterInstallationDto *service.InstallationDto,
) (*EnergyMeterInstallation, error) {
	conn := connect()

	var energyMeterInstallation EnergyMeterInstallation
	result := conn.First(&energyMeterInstallation, energyMeterInstallationDto.GetId())
	return &energyMeterInstallation, result.Error
}

func DeleteEnergyMeterInstallation(
	energyMeterInstallationDto *service.InstallationDto,
) (*EnergyMeterInstallation, error) {
	conn := connect()

	var energyMeterInstallation EnergyMeterInstallation
	result := conn.Delete(&energyMeterInstallation, energyMeterInstallationDto.GetId())
	return &energyMeterInstallation, result.Error
}

func UpdateEnergyMeterInstallation(
	updateEnergyMeterInstallationDto *service.UpdateInstallationDto,
) (*EnergyMeterInstallation, error) {
	conn := connect()

	retirementDate := updateEnergyMeterInstallationDto.GetRetirementDate().AsTime()
	energyMeterInstallation := EnergyMeterInstallation{
		ID:             uint(updateEnergyMeterInstallationDto.GetId()),
		Address:        updateEnergyMeterInstallationDto.GetAddress(),
		RetirementDate: &retirementDate,
		Lines:          uint8(updateEnergyMeterInstallationDto.GetLines()),
		IsActive:       updateEnergyMeterInstallationDto.GetIsActive(),
	}
	result := conn.Save(&energyMeterInstallation)
	return &energyMeterInstallation, result.Error
}
