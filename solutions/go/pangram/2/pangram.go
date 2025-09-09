// Package pangram provides functionality to check if a sentence
// is a pangram, meaning it contains every letter of the English alphabet at least once
package pangram

import (
	"unicode"
)

const numLetters = 26

// IsPangram checks if the input contains every letter of the English alphabet at least once.
func IsPangram(input string) bool {
	letters := make(map[rune]bool, numLetters)
	for _, letter := range input {
		if unicode.IsLetter(letter) {
			letters[unicode.ToLower(letter)] = true
		}
		if len(letters) == 26 {
			return true
		}
	}
	return false
}
