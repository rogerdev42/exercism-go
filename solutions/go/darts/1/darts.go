// Package darts provides scoring logic for a simple darts game.
package darts

const (
	highScoreValue   = 10
	middleScoreValue = 5
	lowScoreValue    = 1

	// Squared distances for each scoring circle
	highScoreDistance   = 1
	middleScoreDistance = 25
	lowScoreDistance    = 100
)

// Score calculates the points for a dart thrown at position (x, y).
func Score(x, y float64) int {
	distance := x*x + y*y
	switch {
	case distance <= highScoreDistance:
		return highScoreValue
	case distance <= middleScoreDistance:
		return middleScoreValue
	case distance <= lowScoreDistance:
		return lowScoreValue
	}
	return 0
}
