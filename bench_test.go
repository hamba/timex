package timex_test

import (
	"testing"
	"time"

	"github.com/hamba/timex"
	"github.com/hamba/timex/mono"
)

func BenchmarkTimexMonoNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = mono.Now()
		}
	})
}

func BenchmarkTimexMonoSince(b *testing.B) {
	start := mono.Now()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = mono.Since(start)
		}
	})
}

func BenchmarkTimexNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = timex.Now()
		}
	})
}

func BenchmarkTimexSince(b *testing.B) {
	start := timex.Now()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = timex.Since(start)
		}
	})
}

func BenchmarkTimexUnix(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = timex.Unix()
		}
	})
}

func BenchmarkTimeNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = time.Now().UnixNano()
		}
	})
}

func BenchmarkTimeSince(b *testing.B) {
	start := time.Now()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = time.Since(start)
		}
	})
}

func BenchmarkTimeUnix(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = time.Now().Unix()
		}
	})
}
