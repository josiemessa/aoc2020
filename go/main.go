package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/josiemessa/aoc2020/go/day3"
)

func main() {
	//fmt.Println("Day 1: Part 1")
	//d1 := day1.Solver{Target: 2020, Input: d1Input}
	//p1 := d1.Solve2()
	//fmt.Println("Solution: ", p1)
	//
	//fmt.Println("Day 1: Part 2")
	//p2 := d1.Solve3()
	//fmt.Println("Solution: ", p2)
	//
	//d2 := day2.PasswordSolver{Input: readFile("C:\\dev\\src\\github.com\\josiemessa\\aoc2020\\inputs\\day2")}
	//fmt.Println("Day 2: Part 1")
	//d2p1 := d2.Solve(true)
	//fmt.Println("Solution ", d2p1)
	//
	//fmt.Println("Day 2: Part 2")
	//d2p2 := d2.Solve(false)
	//fmt.Println("Solution ", d2p2)
	defns := []day3.SlopeDefinition{
		{day3.Right: 1, day3.Down: 1},
		{day3.Right: 3, day3.Down: 1},
		{day3.Right: 5, day3.Down: 1},
		{day3.Right: 7, day3.Down: 1},
		{day3.Right: 1, day3.Down: 2},
	}

	fmt.Println("Day 3: Part 2")
	soln := 1
	for _, defn := range defns {
		d3 := &day3.TobogganSolver{Input: readFile("C:\\dev\\src\\github.com\\josiemessa\\aoc2020\\inputs\\day3"),
			Slope: defn,
		}
		d3p2 := d3.Solve()
		fmt.Println(d3p2)
		soln *= d3p2
	}
	fmt.Println("Final answer:", soln)
}

func readFile(name string) []string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error", err)
	}
	return lines
}
