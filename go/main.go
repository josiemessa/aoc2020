package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/josiemessa/aoc2020/go/day4"
)

func main() {
	f, err := os.Open("C:\\dev\\src\\github.com\\josiemessa\\aoc2020\\inputs\\day4")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()

	//fmt.Println("Part 1")
	//d4p1 := day4.Solve(f, false)
	//fmt.Println(d4p1)

	fmt.Println("Part 2")
	d4p2 := day4.Solve(f, true)
	fmt.Println(d4p2)
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
