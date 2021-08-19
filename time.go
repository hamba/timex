// Package timex provides a fast wall clock source.
package timex

import "time"

// Since is analogous to https://golang.org/pkg/time/#Since.
func Since(s int64) time.Duration {
	return time.Duration(Now() - s)
}
