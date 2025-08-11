package cars

const (
    NumCarsInGroup = 10
	CarPrice = 10000
    GroupCarPrice = 9500
    MinutesInHour = 60
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * successRate / 100 / MinutesInHour)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	numGroups := carsCount / NumCarsInGroup
    numIndCars := carsCount % NumCarsInGroup
    return uint(numGroups * NumCarsInGroup * GroupCarPrice + numIndCars * CarPrice)
}
