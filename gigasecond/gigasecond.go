// Package gigasecond provides a method to calculate the moment
// when someone has lived for 10^9 seconds.
package gigasecond

import "time"

// AddGigasecond adds a gigasecond to a given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1000000000))
}
