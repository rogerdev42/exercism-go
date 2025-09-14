package romannumerals

import (
	"errors"
	"strings"
)

const (
	minNumber = 1
	maxNumber = 3999
)

var dir = map[int]map[int]string{
	1000: {1: "M"},
	100:  {1: "C", 2: "D", 3: "M"},
	10:   {1: "X", 2: "L", 3: "C"},
	1:    {1: "I", 2: "V", 3: "X"},
}

var bases = [...]int{1000, 100, 10, 1}

func ToRomanNumeral(input int) (string, error) {
	if input < minNumber || input > maxNumber {
		return "", errors.New("input must be between 1 and 3999")
	}

	var b strings.Builder
	for _, base := range bases {
		digit := input / base
		if digit != 0 {
			part, _ := romanNumeral(base, digit)
			b.WriteString(part)
			input %= base
		}
	}
	return b.String(), nil
}

func romanNumeral(base, digit int) (string, error) {
	switch digit {
	case 1, 2, 3:
		return strings.Repeat(dir[base][1], digit), nil
	case 4:
		return dir[base][1] + dir[base][2], nil
	case 5:
		return dir[base][2], nil
	case 6, 7, 8:
		return dir[base][2] + strings.Repeat(dir[base][1], digit-5), nil
	case 9:
		return dir[base][1] + dir[base][3], nil
	default:
		return "", errors.New("digit must be between 1 and 9")
	}
}
