package day10

import (
	"sort"

	"github.com/josiemessa/aoc2020/go/utils"
)

func SolveP1(input []string) int {
	diffs := map[int]int{
		1: 0, 2: 0, 3: 0,
	}
	joltConverters := utils.SliceAtoi(input)
	sort.Ints(joltConverters)

	diff := joltConverters[0]
	diffs[diff]++
	for i, joltRating := range joltConverters {
		if i == len(joltConverters)-1 {
			diff = 3
		} else {
			diff = joltConverters[i+1] - joltRating
		}
		diffs[diff]++
	}
	return diffs[1] * diffs[3]
}

func SolveP2(input []string) int {
	joltConverters := utils.SliceAtoi(input)
	joltConverters = append(joltConverters, 0)
	sort.Ints(joltConverters)

	// for the joltConverter in the phone (highest joltRating+3) there is only one path
	highest := joltConverters[len(joltConverters)-1]

	// how many paths to destination from the current node (key)
	graph := map[int]int{highest: 1}

	// from the highest number is the converter chain, calculate how many paths lead into this node
	for i := len(joltConverters) - 1; i > 0; i-- {
		currentNode := joltConverters[i]
		//fmt.Println("\nCurrent Node:", currentNode)
		pathsIn := 0
		for j := i - 1; j >= 0; j-- {
			potential := joltConverters[j]
			if currentNode-potential > 3 {
				// not a valid edge in this graph
				break
			}
			//fmt.Print(" ",potential)
			pathsIn++
			// this is a path into the current node
			_, ok := graph[potential]
			if !ok {
				// new path
				graph[potential] = graph[currentNode]
			} else {
				graph[potential] += graph[currentNode]
			}
		}
		//fmt.Println("\nGraph:", graph)
	}
	return graph[0]
}
