// Package gigasecond provides a function to compute the moment
// one gigasecond (10^9 seconds) after a given time.
package gigasecond

import "time"

// AddGigasecond returns the time that is exactly one gigasecond (1,000,000,000 seconds)
// later than the given input time t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
