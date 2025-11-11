package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	rate := float64(productionRate/60) * (successRate / 100)
	return int(rate)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	totalCost := uint(0)
	remainingCars := carsCount

	if carsCount == 0 {
		return uint(0)
	}

	for i := carsCount; i >= 10; i -= 10 {
		totalCost += 95000
		remainingCars -= 10
	}

	if remainingCars > 0 {
		totalCost += 10000 * uint(remainingCars)
	}

	return totalCost
}
