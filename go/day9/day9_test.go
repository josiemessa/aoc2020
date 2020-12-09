package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	input := []string{
		"35",
		"20",
		"15",
		"25",
		"47",
		"40",
		"62",
		"55",
		"65",
		"95",
		"102",
		"117",
		"150",
		"182",
		"127",
		"219",
		"299",
		"277",
		"309",
		"576",
	}
	require.Equal(t, 127, Solve(input, 5, true))
	require.Equal(t, 62, Solve(input, 5, false))
}
