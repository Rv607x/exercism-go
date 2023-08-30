package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var TotalCarsPerHour = float64(productionRate) * (successRate / 100)
	return int(TotalCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var carsfloat = float64(carsCount / 10)
	var ipart = int(carsfloat)
	var decimalpart = carsCount % 10
	return uint((ipart * 95000) + (decimalpart * 10000))
}
