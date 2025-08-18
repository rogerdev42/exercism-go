// Package isogram checks if a word or phrase is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram reports whether word contains no repeating letters.
// Non-letter characters are ignored.
func IsIsogram(word string) bool {
	seen := make(map[rune]struct{}, len(word))
	for _, r := range strings.ToLower(word) {
		if !unicode.IsLetter(r) {
			continue
		}
		if _, ok := seen[r]; ok {
			return false
		}
		seen[r] = struct{}{}
	}
	return true
}
