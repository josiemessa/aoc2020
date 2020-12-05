package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/josiemessa/aoc2020/go/day5"
)

func main() {
	input := readFile("C:\\dev\\src\\github.com\\josiemessa\\aoc2020\\inputs\\day5")

	//fmt.Println("Day 5 Part 1")
	//d5p2 := day5.SolveP1(input)
	//fmt.Println(d5p2)

	fmt.Println("Day 5 Part 2")
	d5p2 := day5.SolveP2(input)
	fmt.Println(d5p2)
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
