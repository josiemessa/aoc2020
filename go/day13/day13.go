package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SolveP1(input []string) int {
	timestamp, busIDs := parse(input)
	// put an arbitrary break here so this doesn't spin forever

	for i := timestamp; i < timestamp*2; i++ {
		for _, id := range busIDs {
			x := i % id
			if x == 0 {
				return id * (i - timestamp)
			}
		}
	}
	return 0
}

func bruteForceP2(input []string) int {
	// x === 0 (mod 7)
	// x === 1 (mod 13)
	// x === 4 (mod 59)
	// x === 6 (mod 31)
	// x === 7 (mod 19)

	busIDs := parseP2(input)
	product := 1
	for _, id := range busIDs {
		product *= id
	}

	for y := 1; y < product; y++ {
		x := busIDs[0] * y
		found := true
		for offset, id := range busIDs {
			if (x+offset)%id != 0 {
				found = false
			}
		}
		if found {
			return x
		}
	}
	return 0
}

func SolveP2(input []string) int {
	busIDs := parseP2(input)
	time := busIDs[0]
	interval := busIDs[0]

	for offset, id := range busIDs {
		if offset == 0 {
			continue
		}
		for {
			if (time+offset)%id != 0 {
				time += interval
				continue
			}
			fmt.Printf("ID: %d, offset: %d, first sync: %d\n", id, offset, time)
			interval *= id
			break
		}
	}

	return time
}

func parseP2(input []string) map[int]int {
	busIDsStr := strings.Split(input[1], ",")
	busIDs := make(map[int]int)
	for i, idStr := range busIDsStr {
		if idStr == "x" {
			continue
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("Could not parse bus ID %q, on line %q", idStr, input[1])
		}
		busIDs[i] = id
	}
	return busIDs
}

func parse(input []string) (int, []int) {
	timestamp, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatalf("Could not parse current timestamp on line %q", input[0])
	}

	busIDsStr := strings.Split(input[1], ",")
	var busIDs []int
	for _, idStr := range busIDsStr {
		if idStr == "x" {
			continue
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("Could not parse bus ID %q, on line %q", idStr, input[1])
		}
		busIDs = append(busIDs, id)
	}
	return timestamp, busIDs
}
