package database

import "time"

type EnergyMeter struct {
	Serial             string `gorm:"primaryKey"`
	EnergyMeterBrandID uint

	EnergyMeterInstallation []EnergyMeterInstallation
}

type EnergyMeterBrand struct {
	ID   uint `gorm:"primaryKey"`
	Name string

	EnergyMeter []EnergyMeter
}

type EnergyMeterInstallation struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	Address          string
	InstallationDate *time.Time
	RetirementDate   *time.Time
	Lines            uint8
	IsActive         bool
	CreatedAt        *time.Time

	EnergyMeterSerial string
}
