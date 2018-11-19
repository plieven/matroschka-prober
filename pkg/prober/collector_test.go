package prober

import (
	"testing"
	"time"

	"github.com/exaring/matroschka-prober/pkg/config"
	"github.com/stretchr/testify/assert"
)

func uint64ptr(v uint64) *uint64 {
	return &v
}

type mockClock struct {
	t time.Time
}

func (m mockClock) Now() time.Time {
	return m.t
}

func TestLastFinishedMeasurement(t *testing.T) {
	tests := []struct {
		name     string
		p        *Prober
		expected int64
	}{
		{
			name: "Test #1",
			p: &Prober{
				clock: mockClock{
					t: time.Unix(1542556558, 0),
				},
				path: config.Path{
					MeasurementLengthMS: uint64ptr(1000),
					TimeoutMS:           uint64ptr(200),
				},
			},
			expected: 1542556556,
		},
		{
			name: "Test #2",
			p: &Prober{
				clock: mockClock{
					t: time.Unix(1542556558, 250000000),
				},
				path: config.Path{
					MeasurementLengthMS: uint64ptr(1000),
					TimeoutMS:           uint64ptr(200),
				},
			},
			expected: 1542556557,
		},
	}

	for _, test := range tests {
		ts := test.p.lastFinishedMeasurement()
		assert.Equalf(t, test.expected, ts, test.name)
	}
}