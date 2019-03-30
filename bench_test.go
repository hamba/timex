package timex_test

import (
	"testing"
	"time"

	"github.com/hamba/timex"
	"github.com/hamba/timex/mono"
)

func BenchmarkMonoNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = mono.Now()
	}
}

func BenchmarkMonoSince(b *testing.B) {
	start := mono.Now()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mono.Since(start)
	}
}

func BenchmarkTimexNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = timex.Now()
	}
}

func BenchmarkTimexSince(b *testing.B) {
	start := timex.Now()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = timex.Since(start)
	}
}

func BenchmarkTimexUnix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = timex.Unix()
	}
}

func BenchmarkWallTimeNow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = time.Now().UnixNano()
	}
}

func BenchmarkWallTimeSince(b *testing.B) {
	start := time.Now()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = time.Since(start)
	}
}
