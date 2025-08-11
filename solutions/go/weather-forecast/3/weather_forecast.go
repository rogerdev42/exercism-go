// Package weather provides a weather forecasting tool.
package weather

var (
    // currentCondition represents a weather condition.
    currentCondition string
    // currentLocation represents a city.
    currentLocation string
)

// Forecast returns a weather forecast. Takes two arguments: city and weather condition.
func Forecast(city, condition string) string {
	currentLocation, currentCondition = city, condition
	return currentLocation + " - current weather condition: " + currentCondition
}
