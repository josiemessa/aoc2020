package main

import "fmt"

var turns = make(map[int]int) // the key is a number that has been spoken, the value is the last turn that was spoken on

func main() {
	start := []int{9, 3, 1, 0, 8, 4}
	fmt.Println(Solve(start, 30000000))
}

func Solve(start []int, max int) int {
	turns = make(map[int]int)
	// Part 1
	// populate turn map with starting numbers
	startLength := len(start)
	for i, num := range start[0 : startLength-1] {
		turns[num] = i + 1
	}

	var nextTurn int
	currentTurn := start[startLength-1] // last number of the starting numbers
	for i := startLength; i < max; i++ {
		nextTurn = calcNextTurn(currentTurn, i)
		turns[currentTurn] = i
		currentTurn = nextTurn
	}
	return currentTurn
}

// * If that was the first time the number has been spoken, the current player says 0.
// * Otherwise, the number had been spoken before; the current player announces
//   how many turns apart the number is from when it was previously spoken.
func calcNextTurn(num int, turn int) int {
	lastTurn, ok := turns[num]
	if !ok {
		return 0
	}

	// calculate the age since the last time we heard this spoken
	return turn - lastTurn

}
