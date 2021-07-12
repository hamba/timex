// Package timex provides a fast wall clock source.
package timex

import (
	"time"
	// Required in order to import walltime.
	_ "unsafe"
)

//go:linkname walltime runtime.walltime
func walltime() (int64, int32)

// Now returns the current time in nanoseconds.
func Now() int64 {
	sec, nsec := walltime()
	return sec*1000000000 + int64(nsec)
}

// Since is analogous to https://golang.org/pkg/time/#Since.
func Since(s int64) time.Duration {
	return time.Duration(Now() - s)
}

// Unix returns the current time in seconds.
func Unix() int64 {
	sec, _ := walltime()
	return sec
}
