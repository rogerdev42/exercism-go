// Package etl converts grouped letter scores to a flat mapping.
package etl

import (
	"strings"
)

// Transform converts score->letters to letter->score.
func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
	for i, v := range in {
		for _, w := range v {
			result[strings.ToLower(w)] = i
		}
	}
	return result
}
