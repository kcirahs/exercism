// Package diffsquares provides utility functions to find the difference between the
// square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSums calculates the square of the sum of the first N natural numbers.
func SquareOfSums(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	sqOfSum := sum * sum
	return sqOfSum
}

// SumOfSquares calculates the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int {
	var sumOfSq int
	for i := 1; i <= n; i++ {
		sumOfSq = sumOfSq + i*i
	}
	return sumOfSq
}

// Difference calculates the difference between the square of the sum and
// the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
