// Package weather provides a weather forecasting tool.
package weather

// CurrentCondition represents a weather condition.
var CurrentCondition string
// CurrentLocation represents a city.
var CurrentLocation string

// Forecast returns a weather forecast. Takes two arguments: city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
