package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/josiemessa/aoc2020/go/day7"
)

func main() {
	input := readFileAsText("C:\\dev\\src\\github.com\\josiemessa\\aoc2020\\inputs\\day7")
	fmt.Println("Day 7")
	d7p2 := day7.Solve("shiny gold", input, true)
	fmt.Println("Part 2:", d7p2)
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
