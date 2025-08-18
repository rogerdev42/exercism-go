// Package diffsquares provides functions to calculate
// the square of sum, sum of squares, and their difference.
package diffsquares

// SquareOfSum returns (1 + 2 + ... + n)².
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns 1² + 2² + ... + n².
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns SquareOfSum(n) - SumOfSquares(n).
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
