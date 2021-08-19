//go:build !go1.17 || !(linux && amd64)
// +build !go1.17 !linux !amd64

package timex

import _ "unsafe" // for go:linkname

//go:linkname walltime runtime.walltime
func walltime() (int64, int32)

// Now returns the current time in nanoseconds.
func Now() int64 {
	sec, nsec := walltime()
	return sec*1000000000 + int64(nsec)
}

// Unix returns the current time in seconds.
func Unix() int64 {
	sec, _ := walltime()
	return sec
}
