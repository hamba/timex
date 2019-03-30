package mono_test

import (
	"testing"
	"time"

	"github.com/hamba/timex/mono"
	"github.com/stretchr/testify/assert"
)

func TestSince(t *testing.T) {
	then := mono.Now() - 1000

	d := mono.Since(then)

	assert.True(t, time.Duration(1000) <= d)
}
