// Package bob returns the answers that Bob gives to various remarks.
package bob

import (
	"strings"
	"unicode"
)

// Hey processes the remark and returns the appropriate answer.
func Hey(remark string) string {
	//trim whitespace from beginning and end of string
	remark = strings.TrimSpace(remark)
	//do checks for the 3 conditions and store results as boolean variables
	//check for silence
	silence := remark == ""
	//check for question
	question := strings.HasSuffix(remark, "?")
	//check for yelling
	yell := false
	for _, r := range remark {
		//skips input that is not a letter
		//returns false and stops checking if there is lowercase
		//returns true if all letters checked are uppercase
		if unicode.IsLetter(r) == false {
			continue
		} else if unicode.IsLower(r) == true {
			yell = false
			break
		} else if unicode.IsUpper(r) == true {
			yell = true
		}
	}
	// process boolean variables to give correct responses
	if silence {
		return "Fine. Be that way!"
	} else if question && yell {
		return "Calm down, I know what I'm doing!"
	} else if question {
		return "Sure."
	} else if yell {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}

}
