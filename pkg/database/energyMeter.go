package database

func CreateEnergyMeterBrand(name string) error {
	conn := connect()

    result := conn.Create(&EnergyMeterBrand{
		Name: name,
	})

    return result.Error
}

func CreateEnergyMeter() {
	conn := connect()

	conn.Create(&EnergyMeter{})
}
