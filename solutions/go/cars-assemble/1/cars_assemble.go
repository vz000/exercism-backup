package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarsPerHour float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
    var carsPerMinute int = int(workingCarsPerHour/60)
    return carsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const tenCarsPrice = 95000
    const individualCarsPrice = 10000
	var groupsOfTen int = int(carsCount/10)
    groupsOfTenTotal := groupsOfTen * tenCarsPrice
    var individualCars int = carsCount - (groupsOfTen * 10)
    individualTotal := individualCars * individualCarsPrice
    return uint(groupsOfTenTotal + individualTotal)
}
