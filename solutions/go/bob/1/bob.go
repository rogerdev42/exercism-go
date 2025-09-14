package bob

import (
	"strings"
	"unicode"
)

// Hey analyzes the remark and returns Bob's response.
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)

	if trimmed == "" {
		return "Fine. Be that way!"
	}

	hasLetters := false
	isYelling := true
	for _, r := range trimmed {
		if unicode.IsLetter(r) {
			hasLetters = true
			if !unicode.IsUpper(r) {
				isYelling = false
			}
		}
	}
	if !hasLetters {
		isYelling = false
	}

	isQuestion := strings.HasSuffix(trimmed, "?")

	switch {
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
