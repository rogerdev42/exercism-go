// Package dna provides functions to analyze DNA sequences.
package dna

import "fmt"

// Histogram maps nucleotides to their count.
type Histogram map[int32]int

// DNA represents a DNA strand.
type DNA string

// Counts returns a histogram of the nucleotides in the DNA strand.
// It returns an error if the strand contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, b := range d {
		if _, ok := h[b]; !ok {
			return h, fmt.Errorf("invalid DNA: %c", b)
		}
		h[b]++
	}
	return h, nil
}
