// Package collatzconjecture implements step counting for the Collatz sequence.
package collatzconjecture

import (
	"errors"
)

// CollatzConjecture returns the number of steps to reach 1 in the Collatz sequence for n.
// Returns an error if n <= 0.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be a positive integer")
	}
	counter := 0
	for n > 1 {
		switch {
		case n%2 == 0:
			n /= 2
		default:
			n = 3*n + 1
		}
		counter++
	}
	return counter, nil
}
