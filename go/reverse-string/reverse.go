// Package reverse provides a function to reverse a string.
package reverse

// String takes a string and reverses the order of the individual letters.
func String(input string) string {
	lenInput := len(input)
	inputSlice := []byte(input)
	for i, j := 0, lenInput-1; i < j; i, j = i+1, j-1 {
		inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
	}
	reversed := string(inputSlice)
	return reversed
}
