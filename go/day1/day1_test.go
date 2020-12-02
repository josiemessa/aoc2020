package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	// Example inputs by AOC
	exampleInputs = []int{1721,
		979,
		366,
		299,
		675,
		1456}
)

func TestSolver_Solve2(t *testing.T) {
	s := &Solver{
		Input:  exampleInputs,
		Target: 2020,
	}

	require.Equal(t, 514579, s.Solve2())
}

func TestSolver_findPair(t *testing.T) {
	s := &Solver{
		Input:  exampleInputs,
		Target: 2020,
	}
	actual1, actual2, ok := s.findPair()
	require.True(t, ok)
	require.Equal(t, 299, actual1)
	require.Equal(t, 1721, actual2)
}

func TestSolver_Solve3(t *testing.T) {
	s := &Solver{Input: exampleInputs, Target: 2020}

	require.Equal(t, 241861950, s.Solve3())
}

func TestSolver_FindTriple(t *testing.T) {
	s := &Solver{Input: exampleInputs, Target: 2020}

	actual, ok := s.findTriple()
	require.True(t, ok)
	require.Len(t, actual, 3)
	//979, 366, and 675
	require.Contains(t, actual, 979)
	require.Contains(t, actual, 366)
	require.Contains(t, actual, 675)
}
