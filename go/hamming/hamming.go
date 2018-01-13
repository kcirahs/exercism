// Package hamming provides functions for comparing strands of DNA.
package hamming

import "errors"

// Distance takes 2 strings representing DNA strands and compares them,
// returning the number of differences and any errors.
func Distance(a, b string) (int, error) {
	//checks for equal length DNA sequences
	if len(a) != len(b) {
		return -1, errors.New("lengths do not match")
	}
	//counts the differences
	n := 0
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n, nil
}
