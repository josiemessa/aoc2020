package day8

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func SolveP1(program []string) int {
	c := InitComputer()
	ExecuteProgram(c, program)
	return c.Vars["acc"]
}

func InitComputer() *Computer {
	c := NewComputer(map[string]int{"acc": 0})
	c.RegisterFunction("acc", func(i int) error {
		c.Vars["acc"] += i
		c.CurrentLine++
		return nil
	})
	c.RegisterFunction("jmp", func(i int) error {
		c.CurrentLine += i
		return nil
	})
	c.RegisterFunction("nop", func(i int) error {
		c.CurrentLine++
		return nil
	})
	return c
}

func ExecuteProgram(c *Computer, program []string) map[int]struct{} {
	linesExecuted := make(map[int]struct{})
	c.CurrentLine = 0
	for c.CurrentLine < len(program) {
		// We've seen this line before, which means we're about to recurse, exit
		if _, ok := linesExecuted[c.CurrentLine]; ok {
			return linesExecuted
		}
		linesExecuted[c.CurrentLine] = struct{}{}
		if err := c.ExecuteProgramLine(program[c.CurrentLine]); err != nil {
			log.Fatalln(err)
		}
	}
	c.CurrentLine = len(program)
	return linesExecuted
}

func SolveP2(input []string) int {
	c := InitComputer()
	// run through the original program. We know the problem exists in the lines that are executed so limit search to that.
	linesExecuted := ExecuteProgram(c, input)
	var lineSlice []int
	for lineNo := range linesExecuted {
		lineSlice = append(lineSlice, lineNo)
	}
	sort.Ints(lineSlice)

	// iterate over lines executed and try changing them
	for _, lineNo := range lineSlice {
		// there should only be one operation per line so this should work
		var replace string
		if strings.Contains(input[lineNo], "jmp") {
			replace = strings.Replace(input[lineNo], "jmp", "nop", -1)
		} else {
			replace = strings.Replace(input[lineNo], "nop", "jmp", -1)
		}

		modInput := copyAndInsert(input, lineNo, replace)
		newC := InitComputer()
		ExecuteProgram(newC, modInput)
		fmt.Println(modInput, newC.Vars)
		if newC.CurrentLine == len(input) {
			fmt.Println(modInput)
			return newC.Vars["acc"]
		}
	}
	return 0
}

func copyAndInsert(src []string, index int, val string) []string {
	var result []string
	for i, s := range src {
		if i == index {
			result = append(result, val)
		} else {
			result = append(result, s)
		}
	}
	return result
}
