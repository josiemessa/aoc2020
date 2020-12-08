package day8

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	inputString := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	input := strings.Split(inputString, "\n")
	require.Equal(t, 5, SolveP1(input))
	require.EqualValues(t, 8, SolveP2(input))
}
