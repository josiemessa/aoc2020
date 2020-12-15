package day14

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type computerP1 struct {
	mem  map[string]int64
	mask string
}

type computerP2 struct {
	mem  map[int64]int64
	mask string
}

type instructionP1 struct {
	mask  string
	store map[string]int64
}

type instructionP2 struct {
	mask  string
	store map[int64]int64
}

func SolveP1(input []string) int64 {
	c := &computerP1{
		mem: make(map[string]int64),
	}
	for _, line := range input {
		inst := parsep1(line)
		if inst.store != nil {
			// this will only have one entry
			for addr, val := range inst.store {
				c.mem[addr] = c.applyMaskToValue(val)
			}
		} else {
			c.mask = inst.mask
		}
	}

	var result int64
	for _, val := range c.mem {
		result += val
	}

	return result
}

func SolveP2(input []string) int64 {
	c := &computerP2{
		mem: make(map[int64]int64),
	}
	for _, line := range input {
		inst := parseP2(line)
		if inst.store != nil {
			// this will only have one entry
			for addr, val := range inst.store {
				c.applyMaskToAddr(addr, val)
			}
		} else {
			c.mask = inst.mask
		}
	}

	var result int64
	for _, val := range c.mem {
		result += val
	}

	return result
}

func (c *computerP2) applyMaskToAddr(addr, val int64) {
	binary := strconv.FormatInt(addr, 2)
	// use printf to pad binary representation with zeroes
	binary = fmt.Sprintf("%0*s", len(c.mask), binary)

	// build up new value by replacing binary rep with mask values
	addrs := []string{""}
	for i, char := range c.mask {
		if char == x {
			var floating []string
			for j, newAddr := range addrs {
				addrs[j] = newAddr + "0"
				floating = append(floating, newAddr+"1")
			}
			addrs = append(addrs, floating...)
			continue
		}
		if char == '1' {
			for j, newAddr := range addrs {
				addrs[j] = newAddr + "1"
			}
			continue
		}
		for j, newAddr := range addrs {
			addrs[j] = newAddr + string(binary[i])
		}
	}

	for _, newAddr := range addrs {
		i, err := strconv.ParseInt(newAddr, 2, 64)
		if err != nil {
			log.Fatalf("could not convert masked value back to int. New val:%q, old val: %d, mask: %q with erorr %q",
				newAddr, addr, c.mask, err)
		}
		c.mem[i] = val
	}

	return
}

const x = 'X'

func (c *computerP1) applyMaskToValue(val int64) int64 {
	binary := strconv.FormatInt(val, 2)
	// use printf to pad binary representation with zeroes
	binary = fmt.Sprintf("%0*s", len(c.mask), binary)

	// build up new value by replacing binary rep with mask values
	newVal := ""
	for i, char := range c.mask {
		if char == x {
			newVal += string(binary[i])
			continue
		}
		newVal += string(c.mask[i])
	}
	i, err := strconv.ParseInt(newVal, 2, 64)
	if err != nil {
		log.Fatalf("could not convert masked value back to int. New val:%q, old val: %d, mask: %q with erorr %q",
			newVal, val, c.mask, err)
	}
	return i
}

func parsep1(line string) (inst instructionP1) {
	kv := strings.Split(line, " = ")
	if kv[0] == "mask" {
		inst.mask = kv[1]
		return
	}
	if kv[0][:3] != "mem" {
		log.Fatalf("failed to parse line %q", line)
	}

	// grab memory address
	addr := strings.TrimSuffix(strings.TrimPrefix(kv[0][3:], "["), "]")
	i, err := strconv.ParseInt(kv[1], 10, 64)
	if err != nil {
		log.Fatalf("failed to parse line %q with err %#v", line, err)
	}
	inst.store = map[string]int64{addr: i}
	return
}

func parseP2(line string) (inst instructionP2) {
	kv := strings.Split(line, " = ")
	if kv[0] == "mask" {
		inst.mask = kv[1]
		return
	}
	if kv[0][:3] != "mem" {
		log.Fatalf("failed to parse line %q", line)
	}

	// grab memory address
	addrString := strings.TrimSuffix(strings.TrimPrefix(kv[0][3:], "["), "]")
	i, err := strconv.ParseInt(kv[1], 10, 64)
	if err != nil {
		log.Fatalf("failed to parse line %q with err %#v", line, err)
	}
	addr, err2 := strconv.ParseInt(addrString, 10, 64)
	if err2 != nil {
		log.Fatalf("failed to parse line %q with err %#v", line, err)
	}
	inst.store = map[int64]int64{addr: i}
	return
}
