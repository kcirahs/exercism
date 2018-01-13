// Package raindrops provides functions for translating numbers into
// raindrop Speak.
package raindrops

import "strconv"

type convTable []struct {
	factor int
	speak  string
}

var raindropSpeak = convTable{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert takes an int and compares its factors to the conversion table
// and returns the string of translated words.
func Convert(a int) string {

	var rainSaysWhat string
	for i := range raindropSpeak {
		if a%raindropSpeak[i].factor == 0 {
			rainSaysWhat = rainSaysWhat + raindropSpeak[i].speak
		}
	}
	// If no matching factors are found, return the int as a string.
	if rainSaysWhat == "" {
		return strconv.Itoa(a)
	}

	return rainSaysWhat
}
