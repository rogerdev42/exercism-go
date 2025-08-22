// Package space provides a function to calculate age on different planets.
package space

// EarthYearSeconds is the number of seconds in one Earth year (365.25 days).
const EarthYearSeconds = 31557600 // 365.25 * 24 * 60 * 60

// Planet represents the name of a planet.
type Planet string

// orbitalPeriods maps planets to their orbital period relative to an Earth year.
var orbitalPeriods = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns how many years old someone would be on a given planet,
// based on a provided age in seconds.
func Age(seconds float64, planet Planet) float64 {
	period, ok := orbitalPeriods[planet]
	if !ok {
		return -1
	}
	return seconds / (EarthYearSeconds * period)
}
