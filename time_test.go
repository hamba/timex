package timex_test

import (
	"testing"
	"time"

	"github.com/hamba/timex"
	"github.com/stretchr/testify/assert"
)

func TestSince(t *testing.T) {
	then := timex.Now() - 1000

	d := timex.Since(then)

	assert.True(t, time.Duration(1000) <= d)
}

func TestUnix(t *testing.T) {
	want := time.Now().Unix()

	got := timex.Unix()

	assert.Equal(t, want, got)
}
