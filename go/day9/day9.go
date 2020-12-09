package day9

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Solve(input []string, length int, part1 bool) int {
	ints := sliceAtoi(input)
	index, key := solvePart1(ints, length)
	if part1 {
		return key
	}

	// we know our contiguous block has to be in the numbers up to index so only work with that.
	ints = ints[:index]
	block := findContiguousBlock(ints, key)
	fmt.Println(block)

	var (
		min = float64(key) // set to a high number the loop below works, we know key is over the max it could be
		max float64
	)
	for _, i := range block {
		min = math.Min(min, float64(i))
		max = math.Max(max, float64(i))
	}
	return int(min + max)
}

func sum(ints []int) (res int) {
	for _, i := range ints {
		res += i
	}
	return
}

// find a contiguous block of elements in ints which sums to key
func findContiguousBlock(ints []int, key int) (block []int) {
	// Start from the end of the array (as it's roughly increasing) and work backwards
	for i := len(ints) - 1; i >= 0; i-- {
		if len(block) == 0 {
			// start building up a new contiguous block
			block = ints[i : i+1]
		} else {
			// i is decreasing through the indices, so let j=i+len(block),
			// Then j is index of the first valid number we added to this block
			// so we need to move our block slice one index down, which gives us the slice below
			block = ints[i : i+len(block)+1]
		}

		// Even if block has 1 entry this is still useful
		x := sum(block)
		if x == key && len(block) > 1 {
			// we did it!
			break
		}
		if x > key {
			// this block will never sum to key, reset block
			// j = i+len(block) was the first index of this contiguous block, so now start looking from j-1.
			i = i + len(block) - 1
			block = []int{}
		}

	}
	return
}

// Find the first element in ints that is not the sum of any pair of numbers in the previous `length` elements.
// This element that is not a sum is the `key`, and return its index too for later use.
func solvePart1(ints []int, length int) (index, key int) {
	for i, x := range ints {
		if i < length {
			continue
		}
		preamble := ints[i-length : i]
		found := false
		for _, y := range preamble {
			found = found || exists(int(math.Abs(float64(x-y))), preamble)
		}
		if !found {
			index = i
			key = x
			return
		}
	}
	return
}

func sliceAtoi(input []string) []int {
	result := make([]int, len(input))
	for i, s := range input {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not parse line %d: %q", i, s)
		}
		result[i] = x
	}
	return result
}

func exists(i int, ints []int) bool {
	for _, j := range ints {
		if j == i {
			return true
		}
	}
	return false
}
