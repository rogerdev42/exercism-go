package isbn

import (
	"strings"
	"unicode"
)

const isbnLen = 10

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != isbnLen {
		return false
	}
	s := 0
	for i, r := range isbn[:isbnLen-1] {
		if !unicode.IsNumber(r) {
			return false
		}
		s += int(r-'0') * (isbnLen - i)
	}

	last := isbn[isbnLen-1]
	if !unicode.IsNumber(rune(last)) && last != 'X' {
		return false
	}
	m := int(last - '0')
	if last == 'X' {
		m = 10
	}
	s += m
	return s%11 == 0
}
