//go:build go1.17
// +build go1.17

package timex

import _ "unsafe" // for go:linkname

//go:linkname now time.now
func now() (int64, int32, int64)

// Now returns the current time in nanoseconds.
func Now() int64 {
	sec, nsec, _ := now()
	return sec*1000000000 + int64(nsec)
}

// Unix returns the current time in seconds.
func Unix() int64 {
	sec, _, _ := now()
	return sec
}
