// Package hamming implements calculation of the Hamming Distance
// between two equal-length DNA strands.
package hamming

import (
	"errors"
)

// Distance returns the Hamming Distance between strings a and b.
// It counts the number of differing positions.
// Returns an error if the strings are not of equal length.
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("strings must be of equal length")
	}
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
