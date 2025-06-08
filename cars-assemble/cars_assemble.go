package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const CostPerCar uint = 10000
	const GroupOfTenCars uint = 95000
	const CarsPerGroup int = 10

	var numberOfGroups uint = uint(carsCount / CarsPerGroup)
	var remainingCars uint = uint(carsCount % CarsPerGroup)

	return (numberOfGroups * GroupOfTenCars) + (remainingCars * CostPerCar)
}
