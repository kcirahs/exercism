// Package acronym contains the function to abbreviate a long name.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate takes the long name as input and returns the acronym.
func Abbreviate(s string) string {
	//splitfunc separates the string at any non-letter value
	splitfunc := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	sp := strings.FieldsFunc(s, splitfunc)
	//takes the first letter of each word and appends it to a slice
	//of bytes that will be the acronym
	var acrbyte []byte
	for i := range sp {
		v := []byte(sp[i])[0]
		acrbyte = append(acrbyte, v)
	}
	return strings.ToUpper(string(acrbyte))
}
