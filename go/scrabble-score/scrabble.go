// Package scrabble provides utilies for calculating the scrabble
// scores of given words.
package scrabble

import "strings"

// Score checks each letter in the word and adds its value to the
// word score.
func Score(word string) int {

	caps := strings.ToUpper(word)
	value := 0
	for _, key := range caps {
		switch key {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			value++
		case 'D', 'G':
			value += 2
		case 'B', 'C', 'M', 'P':
			value += 3
		case 'F', 'H', 'V', 'W', 'Y':
			value += 4
		case 'K':
			value += 5
		case 'J', 'X':
			value += 8
		case 'Q', 'Z':
			value += 10
		}
	}
	return value
}
