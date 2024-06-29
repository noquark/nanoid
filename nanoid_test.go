package nanoid_test

import (
	"testing"

	"github.com/noquark/nanoid"
)

func TestNew(t *testing.T) {
	t.Run("TestSimple", func(t *testing.T) {
		id, err := nanoid.New()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(id) != 21 {
			t.Errorf("expected ID length of 21, got %d", len(id))
		}
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
					if err == nil {
						t.Errorf("expected error on invalid length, got none")
					}
				} else {
					if err != nil {
						t.Errorf("unexpected error: %v", err)
					}
					if len(id) != tt.expected {
						t.Errorf("expected ID length of %d, got %d", tt.expected, len(id))
					}
				}
			})
		}
	})

	t.Run("TestUnexpected", func(t *testing.T) {
		_, err := nanoid.New(10, 15)
		if err == nil {
			t.Errorf("expected error for unexpected parames, got none")
		}
	})

	t.Run("TestNoCollision", func(t *testing.T) {
		tries := 1000_000
		used := make(map[string]bool)

		for i := 0; i < tries; i++ {
			id := nanoid.Must()
			if used[id] {
				t.Errorf("unexpected collsion")
			}
			used[id] = true
		}
	})
}

func BenchmarkNanoID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = nanoid.New()
	}
}
