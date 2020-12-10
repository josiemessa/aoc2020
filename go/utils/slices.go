package utils

import (
	"log"
	"strconv"
)

func SliceAtoi(input []string) []int {
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

func Exists(i int, ints []int) bool {
	for _, j := range ints {
		if j == i {
			return true
		}
	}
	return false
}

func Sum(ints []int) (res int) {
	for _, i := range ints {
		res += i
	}
	return
}

func Factorial(i int) int {
	res := 1
	for j := 1; j <= i; j++ {
		res *= j
	}
	return res
}

func Product(ints []int, ignoreZeros bool) int {
	res := 1
	for _, i := range ints {
		if ignoreZeros && i == 0 {
			continue
		}
		res *= i
	}
	return res
}
