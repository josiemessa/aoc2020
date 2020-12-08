package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/josiemessa/aoc2020/go/day8"
)

func main() {
	input := readFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2020\\inputs\\day8")
	fmt.Println("Day 8")
	//d8p1 := day8.SolveP1(input)
	//fmt.Println("Part 1:", d8p1)
	d8p2 := day8.SolveP2(input)
	fmt.Println("Part 2:", d8p2)
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
