// Package weather provides a weather forecasting tool.
package weather

var (
    // CurrentCondition represents a weather condition.
    CurrentCondition string
    // CurrentLocation represents a city.
    CurrentLocation string
)

// Forecast returns a weather forecast. Takes two arguments: city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
