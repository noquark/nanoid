package nanoid_test

import (
	"testing"

	"github.com/mohitsinghs/nanoid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("TestSimple", func(t *testing.T) {
		id, err := nanoid.New()
		assert.NoError(t, err, "shouldn't return error")
		assert.Len(t, id, 21, "should return ID of default length")
	})

	t.Run("TestLength", func(t *testing.T) {
		tests := []struct {
			desc      string
			length    int
			expected  int
			shouldErr bool
		}{
			{"TestZero", 0, 0, true},
			{"TestNegative", -1, 0, true},
			{"TestCustom", 7, 7, false},
		}
		for _, tt := range tests {
			t.Run(tt.desc, func(t *testing.T) {
				id, err := nanoid.New(tt.length)
				if tt.shouldErr {
					assert.Error(t, err, "should return error on invalid length")
				} else {
					assert.NoError(t, err, "shouldn't return error")
					assert.Len(t, id, tt.expected, "should return ID of length %d", tt.length)
				}
			})
		}
	})

	t.Run("TestNoCollision", func(t *testing.T) {
		tries := 1000_000
		used := make(map[string]bool)

		for i := 0; i < tries; i++ {
			id := nanoid.Must()
			require.False(t, used[id], "shouldn't return colliding IDs")
			used[id] = true
		}
	})
}

func BenchmarkNanoID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = nanoid.New()
	}
}
