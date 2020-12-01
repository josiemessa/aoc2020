package day1

import (
	"fmt"
	"sort"
)

type Solver struct {
	Input  []int
	Target int
}

func (s *Solver) Solve2() int {
	sort.Ints(s.Input)

	val1, val2, ok := s.findPair()
	if !ok {
		panic("why are you like this?")
	}

	fmt.Println("Two values:", val1, val2)
	return val1 * val2
}

// Find two expense lines from input that sum to the target
// TODO: keep track of pairs that didn't work and somehow remove them from the list
func (s *Solver) findPair() (int, int, bool) {
	iterations := 0

	defer func() {
		fmt.Println("Iterations:", iterations)
	}()

	for i, a := range s.Input {
		for j := i + 1; j < len(s.Input); j++ {
			iterations++
			b := s.Input[j]
			sum := a + b
			if sum == s.Target {
				return a, b, true
			}
			// early break - we've pre-sorted the list, and we're adding up the smallest numbers together first, if we
			// go over the target we're only getting bigger from this point so bomb out early.
			if sum > s.Target {
				break
			}
		}
	}
	return 0, 0, false
}

func (s *Solver) Solve3() int {
	sort.Ints(s.Input)

	result, ok := s.findTriple()
	if !ok {
		panic("why are you always like this?")
	}
	fmt.Println("Three values:", result)
	return result[0] * result[1] * result[2]
}

// Find three expense lines from input that sum to target
func (s *Solver) findTriple() ([]int, bool) {
	iterations := 0
	defer func() { fmt.Println("find triple Iterations:", iterations) }()
	for i, a := range s.Input {
		for j := i + 1; j < len(s.Input); j++ {
			b := s.Input[j]
			sum := a + b
			// early break - we've pre-sorted the list, and we're adding up the smallest numbers together first, if we
			// go over 2020 we're only getting bigger from this point so bomb out early.
			if sum >= 2020 {
				break
			}

			// calculate the remainder
			rem := 2020 - sum
			for k := j + 1; k < len(s.Input); k++ {
				c := s.Input[k]
				iterations++
				if c > rem {
					break
				}
				if c == rem {
					return []int{a, b, c}, true
				}
			}
		}
	}
	return nil, false
}
