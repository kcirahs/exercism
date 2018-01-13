// Package gigasecond uses the "time" standard library to create a function that returns the result of adding 1 gigasecond to a time.Time input.
package gigasecond

import "time"

// This creates a constant of type duration to be used with time package has a value of 1 gigasecond.
const gigasecond = 1e9 * time.Second

// AddGigasecond returns a given time plus a gigasecond.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(gigasecond)
	return t
}
