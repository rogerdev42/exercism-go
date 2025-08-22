// Package luhn implements validation of strings using the Luhn algorithm.
package luhn

import "strings"

// Valid reports whether id is valid according to the Luhn formula.
// Strings shorter than 2 are invalid, spaces are ignored,
// and any non-digit character makes the string invalid.
func Valid(id string) bool {
	s := strings.ReplaceAll(id, " ", "")
	if len(s) < 2 {
		return false
	}

	sum := 0
	parity := len(s) % 2
	for i, r := range s {
		if r < '0' || r > '9' {
			return false
		}
		n := int(r - '0')
		if i%2 == parity {
			n = n*2%10 + n*2/10
		}
		sum += n
	}

	return sum%10 == 0
}
