package day11

import (
	"strings"
)

var (
	dot   = "."[0]
	empty = "L"[0]
	hash  = "#"[0]
)

func Solve(input []string, part1 bool) int {
	for {
		newMap, changes := iterate(input, part1)
		if changes == 0 {
			break
		}
		input = newMap
	}

	hashes := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] == hash {
				hashes++
			}
		}
	}
	return hashes
}

func iterate(seatMap []string, part1 bool) ([]string, int) {
	newSeats := make([]string, len(seatMap))
	var changes int

	for i := range seatMap {
		newSeats[i] = ""
		for j := range strings.TrimSpace(seatMap[i]) {
			if seatMap[i][j] == dot {
				newSeats[i] += string(seatMap[i][j])
				continue
			}

			var adj int
			if part1 {
				adj = calcAdjacent(i, j, seatMap)
			} else {
				adj = calcVisible(i, j, seatMap)
			}

			if seatMap[i][j] == empty && adj == 0 {
				newSeats[i] += string(hash)
				changes++
				continue
			}
			if seatMap[i][j] == hash {
				if part1 && adj >= 4 {
					newSeats[i] += string(empty)
					changes++
					continue
				}
				if !part1 && adj == 5 {
					newSeats[i] += string(empty)
					changes++
					continue
				}
			}
			newSeats[i] += string(seatMap[i][j])
		}
	}
	return newSeats, changes
}

func calcVisible(i, j int, seatMap []string) (visible int) {
	// go right
	visible += findSeatStraight(j+1, // start to right of current element
		func(k int) bool { return k < len(seatMap[i]) }, // row test
		func(k int) int { return k + 1 },                // move right along row
		func(k int) uint8 { return seatMap[i][k] })      // grab element from row

	// go left

	visible += findSeatStraight(j-1, // start to the left of current element
		func(k int) bool { return k >= 0 },         // row test
		func(k int) int { return k - 1 },           // move left along row
		func(k int) uint8 { return seatMap[i][k] }) // grab element from row

	// go up

	visible += findSeatStraight(i-1, // start above current element
		func(k int) bool { return k >= 0 },         // column test
		func(k int) int { return k - 1 },           // move up along column
		func(k int) uint8 { return seatMap[k][j] }) // grab element from column

	// go down
	visible += findSeatStraight(i+1, // start below current element
		func(k int) bool { return k < len(seatMap) }, // column test
		func(k int) int { return k + 1 },             // move down along column
		func(k int) uint8 { return seatMap[k][j] })   // grab element from column

	// go right/up
	visible += findSeatDiag(func() (int, int) { return i - 1, j + 1 },
		func(k, l int) bool { return k >= 0 && l < len(seatMap[k]) },
		func(k, l int) (int, int) { return k - 1, l + 1 },
		func(k, l int) uint8 { return seatMap[k][l] })
	if visible == 5 {
		return
	}

	// go right/down
	visible += findSeatDiag(func() (int, int) { return i + 1, j + 1 },
		func(k, l int) bool { return k < len(seatMap) && l < len(seatMap[k]) },
		func(k, l int) (int, int) { return k + 1, l + 1 },
		func(k, l int) uint8 { return seatMap[k][l] })
	if visible == 5 {
		return
	}

	// go left/down
	visible += findSeatDiag(func() (int, int) { return i + 1, j - 1 },
		func(k, l int) bool { return k < len(seatMap) && l >= 0 },
		func(k, l int) (int, int) { return k + 1, l - 1 },
		func(k, l int) uint8 { return seatMap[k][l] })
	if visible == 5 {
		return
	}

	// go left/up
	visible += findSeatDiag(func() (int, int) { return i - 1, j - 1 },
		func(k, l int) bool { return k >= 0 && l >= 0 },
		func(k, l int) (int, int) { return k - 1, l - 1 },
		func(k, l int) uint8 { return seatMap[k][l] })
	return
}

func findSeatStraight(init int, test func(int) bool, op func(int) int, val func(int) uint8) int {
	for k := init; test(k); k = op(k) {
		seat := val(k)
		switch seat {
		case dot:
			// continue to next iteration of this for loop
			continue
		case hash:
			return 1
		}
		// break out of this for loop
		break
	}
	return 0
}

func findSeatDiag(init func() (int, int), test func(int, int) bool, op func(int, int) (int, int), val func(int, int) uint8) int {
	for k, l := init(); test(k, l); k, l = op(k, l) {
		switch val(k, l) {
		case dot:
			// continue to next iteration of this for loop
			continue
		case hash:
			return 1
		}
		// break out of this for loop
		break
	}
	return 0
}

func calcAdjacent(i, j int, seatMap []string) int {
	var adjacent int

	left := j-1 >= 0
	above := i-1 >= 0
	right := j+1 < len(seatMap[i])
	below := i+1 < len(seatMap)

	if above {
		if seatMap[i-1][j] == hash {
			adjacent++
		}
		if left && seatMap[i-1][j-1] == hash {
			adjacent++
		}
		if right && seatMap[i-1][j+1] == hash {
			adjacent++
		}
	}
	if right && seatMap[i][j+1] == hash {
		adjacent++
	}

	if below {
		if right && seatMap[i+1][j+1] == hash {
			adjacent++
		}
		if seatMap[i+1][j] == hash {
			adjacent++
		}
		if left && seatMap[i+1][j-1] == hash {
			adjacent++
		}
	}
	if left && seatMap[i][j-1] == hash {
		adjacent++
	}
	return adjacent
}
