package day12

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type direction struct {
	x    int
	y    int
	turn map[string]string
}

var (
	directions = map[string]direction{
		"N": {x: 0, y: 1, turn: map[string]string{"R": "E", "L": "W"}},
		"S": {x: 0, y: -1, turn: map[string]string{"R": "W", "L": "E"}},
		"E": {x: 1, y: 0, turn: map[string]string{"R": "S", "L": "N"}},
		"W": {x: -1, y: 0, turn: map[string]string{"R": "N", "L": "S"}},
	}
	forward     = "E"
	waypointPos = [2]int{10, 1} // relative position to the ship
	shipPos     = [2]int{0, 0}
	waypointTun = map[string]func(int, int) (int, int){
		"R": func(x, y int) (int, int) { return y, x * -1 },
		"L": func(x, y int) (int, int) { return y * -1, x },
	}
)

func Solve(input []string, waypoint bool) float64 {
	// reset
	forward = "E"
	waypointPos = [2]int{10, 1} // relative position to the ship
	shipPos = [2]int{0, 0}
	for _, line := range input {
		if waypoint {
			moveWaypoint(line)
		} else {
			moveShip(line)
		}
		fmt.Println(line, shipPos, forward)
	}
	return math.Abs(float64(shipPos[0])) + math.Abs(float64(shipPos[1]))
}

func moveWaypoint(instruction string) {
	moveInstruction := string(instruction[0])
	units, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatalf("Could not parse instruction %q, %#v", instruction, err)
	}
	if dir, ok := directions[moveInstruction]; ok {
		waypointPos[0] += dir.x * units
		waypointPos[1] += dir.y * units
		return
	}
	if moveInstruction == "F" {
		shipPos[0] += waypointPos[0] * units
		shipPos[1] += waypointPos[1] * units
		return
	}
	// must be turning
	for i := units / 90; i > 0; i-- {
		f := waypointTun[moveInstruction]
		waypointPos[0], waypointPos[1] = f(waypointPos[0], waypointPos[1])
	}
}

func moveShip(instruction string) {
	moveInstruction := string(instruction[0])
	units, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatalf("Could not parse instruction %q, %#v", instruction, err)
	}
	if dir, ok := directions[moveInstruction]; ok {
		shipPos[0] += dir.x * units
		shipPos[1] += dir.y * units
		return
	}
	if moveInstruction == "F" {
		shipPos[0] += directions[forward].x * units
		shipPos[1] += directions[forward].y * units
		return
	}
	// must be turning
	for i := units / 90; i > 0; i-- {
		forward = directions[forward].turn[moveInstruction]
	}
}
