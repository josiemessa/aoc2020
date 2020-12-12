package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/josiemessa/aoc2020/go/day12"
)

func main() {
	input := readFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2020\\inputs\\day12")
	fmt.Println("Day 12")
	//d10p1 := day10.SolveP1(input)
	//fmt.Println("Part 1:", d10p1)
	soln := day12.Solve(input)
	fmt.Println("Part 1:", soln)
}

func readFileAsText(name string) string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Could not read file", err)
	}
	return string(contents)
}

func readFileAsLines(name string) []string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()
	fmt.Println(f.Name())

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
