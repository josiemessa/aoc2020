package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#"}

func TestDay2_Part1(t *testing.T) {
	s := &TobogganSolver{Input: input, Slope: SlopeDefinition{Right: 3, Down: 1}}
	require.Equal(t, 7, s.Solve())
}
