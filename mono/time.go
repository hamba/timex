// Package mono provides a fast monotonic clock source.
package mono

import (
	"time"
	_ "unsafe" // Required in order to import nanotime
)

//go:linkname nanotime runtime.nanotime
func nanotime() int64

// Now returns the current time in nanoseconds from a monotonic clock.
//
// The time returned is guaranteed to increase monotonically at a
// constant rate, unlike timex.Now() which is influenced by NTP
// changes.
func Now() int64 {
	return nanotime()
}

// Since is analogous to https://golang.org/pkg/time/#Since.
func Since(s int64) time.Duration {
	return time.Duration(Now() - s)
}
