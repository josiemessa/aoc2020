package day12

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	inputStr := `F10
N3
F7
R90
F11`
	require.Equal(t, 25.0, Solve(strings.Split(inputStr, "\n"), false))
}

func TestSolve2(t *testing.T) {
	inputStr := `F10
N3
F7
R90
F11`
	require.Equal(t, 286.0, Solve(strings.Split(inputStr, "\n"), true))
}
