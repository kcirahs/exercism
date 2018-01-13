//Package isogram provides a function for checking if a string is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram takes a string value and check if any letters repeat, returning
// a boolean value.
func IsIsogram(word string) bool {
	wordLen := len(word)
	wordLower := strings.ToLower(word)
	for i, letter := range wordLower {
		if !unicode.IsLetter(letter) {
			continue
		}
		for j := i + 1; j < wordLen; j++ {
			if wordLower[j] == wordLower[i] {
				return false
			}
		}
	}
	return true
}
