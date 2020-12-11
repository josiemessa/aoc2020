package day11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolveP1(t *testing.T) {
	inputStr := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	require.Equal(t, 37, Solve(strings.Split(inputStr, "\n"), true))
}

func TestFindVisible(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		i        int
		j        int
		expected int
	}{
		{name: "all occupied",
			input: `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`,
			expected: 5, // we break as soon as we hit 5
			i:        4,
			j:        3},
		{
			name: "empty seat blocking",
			input: `.............
.L.L.#.#.#.#.
.............`,
			i:        1,
			j:        1,
			expected: 0,
		},
		{
			name: "all empty",
			input: `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`,
			i:        3,
			j:        3,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, calcVisible(tt.i, tt.j, strings.Split(tt.input, "\n")))
		})
	}
}
