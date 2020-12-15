package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolveP1(t *testing.T) {
	tests := []struct {
		name     string
		start    []int
		expected int
	}{
		{
			name:     "main example",
			start:    []int{0, 3, 6},
			expected: 436,
		},
		{
			name:     "first example",
			start:    []int{1, 3, 2},
			expected: 1,
		},
		{
			name:     "second example",
			start:    []int{2, 1, 3},
			expected: 10,
		},
		{
			name:     "third example",
			start:    []int{1, 2, 3},
			expected: 27,
		},
		{
			name:     "fourth example",
			start:    []int{2, 3, 1},
			expected: 78,
		},
		{
			name:     "fifth example",
			start:    []int{3, 2, 1},
			expected: 438,
		},
		{
			name:     "sixth example",
			start:    []int{3, 1, 2},
			expected: 1836,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, Solve(tt.start, 2020))
		})
	}
}
