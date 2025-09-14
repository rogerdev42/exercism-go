package grains

import "fmt"

// Square returns the number of grains on a specific square.
// Square number must be between 1 and 64.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("square must be between 1 and 64, got %d", number)
	}
	return 1 << (number - 1), nil
}

// Total returns the total number of grains on the board.
func Total() uint64 {
	return 1<<64 - 1
}
