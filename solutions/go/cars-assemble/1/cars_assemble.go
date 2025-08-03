package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	rate := float64(productionRate)
    return rate * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(workingCarsPerHour) / 60
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const individualCost = 10000
    const groupOfTen = 95000

    numberOfGroups := carsCount / 10
    indidualCars := carsCount - (numberOfGroups * 10)

    return uint(numberOfGroups * groupOfTen + indidualCars * individualCost)
}
