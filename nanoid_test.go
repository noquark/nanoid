package nanoid_test

import (
	"testing"

	"github.com/mohitsinghs/nanoid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	id, err := nanoid.New()
	assert.NoError(t, err, "shouldn't return error")
	assert.Len(t, id, 21, "should return ID of default length")
}

func TestNegativeLength(t *testing.T) {
	_, err := nanoid.New(-1)
	assert.Error(t, err, "should return error if the requested ID length is invalid")
}

func TestCustomLength(t *testing.T) {
	id, err := nanoid.New(7)
	assert.NoError(t, err, "shouldn't return error")
	assert.Len(t, id, 7, "should return ID of requested length")
}

func TestNoCollision(t *testing.T) {
	tries := 1000_000
	used := make(map[string]bool)

	for i := 0; i < tries; i++ {
		id := nanoid.Must()
		require.False(t, used[id], "shouldn't return colliding IDs")
		used[id] = true
	}
}

func BenchmarkNanoID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = nanoid.New()
	}
}
