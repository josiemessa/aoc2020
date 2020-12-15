package day13

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolveP1(t *testing.T) {
	inputStr := `939
7,13,x,x,59,x,31,19`
	require.Equal(t, 295, SolveP1(strings.Split(inputStr, "\n")))
}

func TestSolveP2(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:  "two values",
			input: []string{"939", "7,13"},
		},
		{
			name:  "three values",
			input: []string{"939", "7,13,x,x,59"},
		},
		{
			name:  "four values",
			input: []string{"939", "7,13,x,x,59,x,31"},
		},
		{
			name:     "five values",
			input:    []string{"939", "7,13,x,x,59,x,31,19"},
			expected: 1068781,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != 0 {
				require.Equal(t, tt.expected, SolveP2(tt.input))
			} else {
				require.Equal(t, bruteForceP2(tt.input), SolveP2(tt.input))
			}
		})
	}

}
