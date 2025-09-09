// Package reverse provides a function to reverse strings.
package reverse

// Reverse takes a string and returns its characters in reverse order.
func Reverse(input string) string {
	runes := []rune(input)
	reversed := make([]rune, len(runes))
	for i, val := range runes {
		reversed[len(runes)-1-i] = val
	}
	return string(reversed)
}
