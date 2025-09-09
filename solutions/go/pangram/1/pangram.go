package pangram

import (
	"unicode"
)

const numLetters = 26

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
