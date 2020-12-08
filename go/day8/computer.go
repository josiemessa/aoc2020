package day8

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Computer struct {
	operations  map[string]func(int) error
	Vars        map[string]int
	CurrentLine int
}

func NewComputer(vars map[string]int) *Computer {
	return &Computer{Vars: vars}
}

func (c *Computer) RegisterFunction(name string, f func(int) error) {
	if c.operations == nil {
		c.operations = make(map[string]func(int) error)
	}
	c.operations[name] = f
}

func (c *Computer) ExecuteProgramLine(instruction string) error {
	split := strings.Split(instruction, " ")
	j, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatalf("Could not parse signed int in instruction %q. Err: %#v", instruction, err)
	}
	op, ok := c.operations[split[0]]
	if !ok {
		return fmt.Errorf("operation %q not found", split[0])
	}
	return op(j)
}
