package day16

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2020/go/utils"
)

func SolveP1(input []string) int {
	rules, scannedTickets := parse(input)

	// first index in validTickets is the order that this ticket appears in the valid list
	// second index is the order that a field appears in the ticket itself
	// the map is a set of all matching fields for this field
	var validTickets []map[string][]int
	var invalid int
	for _, ticket := range scannedTickets {
		invalidFields, matchedFields := IsValid(ticket, rules)
		invalid += utils.Sum(invalidFields)
		if len(invalidFields) == 0 {
			validTickets = append(validTickets, matchedFields)
		}
	}
	return invalid

	//fmt.Printf("Part 1: %d\n", valid)
	//
	//ticketLength := len(rules)
	//finalFields := make([]string, ticketLength)
	//unfinishedFields := make([]map[string]struct{}, ticketLength)
	//
	//// iterate over each field in the ticket
	//for i := 0; i < ticketLength; i++ {
	//
	//	// for each ticket, grab the list of possible fields for the ith field
	//	var ithField []map[string]struct{}
	//	for _, ticket := range validTickets {
	//		ithField = append(ithField, ticket[i])
	//	}
	//	possibleFields := intersection(ithField...)
	//
	//	if len(possibleFields) == 1 {
	//		for field := range possibleFields {
	//			finalFields[i] = field
	//		}
	//	} else if len(possibleFields) == 0 {
	//		log.Fatalf("No valid field for field %d. Possible fields were: %#v", i, ithField)
	//	} else {
	//		unfinishedFields[i] = possibleFields
	//	}
	//}
	//
	//if len(unfinishedFields) > 0 {
	//	for i := range unfinishedFields {
	//		if unfinishedFields[i] == nil {
	//			continue
	//		}
	//		for f := range unfinishedFields[i] {
	//			if finalFields[f] != ""{
	//				delete(unfinishedFields[i], f)
	//			}
	//		}
	//		rest := append(unfinishedFields[0:i],unfinishedFields[i+1:]...)
	//		exFields := exclusion(unfinishedFields[i], rest...)
	//		if len(exFields) != 1 {
	//			fmt.Printf("exFields no worky. possible fields %#v, i: %d\n", exFields, i)
	//			unfinishedFields[i] = exFields
	//		}
	//		for f := range exFields {
	//			finalFields[i] = f
	//		}
	//	}
	//}
	//
	//fmt.Println(finalFields)
	//
	//myTicket := parseYourTicket(input)
	//solution := 1
	//for i, fieldValue := range myTicket {
	//	if strings.Contains(finalFields[i], "departure") {
	//		solution *= fieldValue
	//	}
	//}
	//return solution

}

func SolveP2(input []string) int {
	rules, scannedTickets := parse(input)
	yourTicket := parseYourTicket(input)

	var validTickets []map[string][]int
	for _, ticket := range scannedTickets {
		invalidFields, matchedFields := IsValid(ticket, rules)
		if len(invalidFields) == 0 {
			validTickets = append(validTickets, matchedFields)
		}
	}

	fieldPositions := make(map[string]int)

	// i is ticket number, fields is a map of field name to all positions that the field could appear on the ticket
	for _, fields := range validTickets {
		for f, possiblePositions := range fields {
			if pos, ok := fieldPositions[f]; !ok {
				// add new field position
				if len(possiblePositions) == 1 {
					fmt.Println("Found a field")
					fieldPositions[f] = possiblePositions[0]
				}
			} else {
				// don't keep this position there
				possiblePositions = []int{pos}

				// remove pos from all other possible positions
				for otherFields, otherPositions := range fields {
					if otherFields == f {
						continue
					}
					if exists, index := utils.Exists(pos, otherPositions); exists {
						fields[otherFields] = append(otherPositions[0:index], otherPositions[index+1:]...)
					}
				}
			}
		}

		// early break as we've now found all the correct positions
		if len(fieldPositions) == len(rules) {
			break
		}
	}

	if len(fieldPositions) != len(rules) {
		fmt.Println("another pass required", len(rules)-len(fieldPositions))
	}

	solution := 1
	for field, pos := range fieldPositions {
		if strings.Contains(field, "departure") {
			solution *= yourTicket[pos]
		}
	}
	return solution
}

func intersection(sets ...map[string]struct{}) map[string]struct{} {
	result := make(map[string]struct{})
	for field := range sets[0] {
		result[field] = struct{}{}
	}
	for _, set := range sets {
		for field := range result {
			if _, ok := set[field]; !ok {
				delete(result, field)
			}
		}
	}
	return result
}

func exclusion(a map[string]struct{}, sets ...map[string]struct{}) map[string]struct{} {
	ex := make(map[string]struct{})
	for f := range a {
		ex[f] = struct{}{}
	}
	for _, set := range sets {
		for f := range a {
			if _, ok := set[f]; ok {
				delete(ex, f)
			}
		}
	}
	return ex
}

// in matchedfields, index is order that we're going through the ticket numbers,
// map[string] is the set of matched fields for this ticket
func IsValid(ticket []int, rules map[string][]int) (invalidFields []int, matchedFields map[string][]int) {
	// for each number in ticket, find out which rules it satisfies.
	// * If no rules, then we can immediately return false

	matchedFields = make(map[string][]int)
	for i, num := range ticket {
		valid := true
		for field, rule := range rules {
			if (rule[0] <= num && num <= rule[1]) || (rule[2] <= num && num <= rule[3]) {
				if matchedFields[field] == nil {
					matchedFields[field] = make([]int, 0)
				}
				matchedFields[field] = append(matchedFields[field], i)
			} else {
				valid = false
			}
		}
		if !valid {
			invalidFields = append(invalidFields, num)
		}
	}
	return
}

func parseYourTicket(input []string) []int {
	// scan over tickets
	yourTicketSection := false
	for _, line := range input {
		if strings.Contains(line, "your ticket:") {
			yourTicketSection = true
			continue
		}
		if !yourTicketSection {
			continue
		}

		var ticketFields []int
		split := strings.Split(line, ",")
		for _, a := range split {
			i, err := strconv.Atoi(a)
			if err != nil {
				log.Fatalf("Failed to parse line %q, %#v", line, err)
			}
			ticketFields = append(ticketFields, i)
		}
		return ticketFields
	}
	return nil
}

func parse(input []string) (map[string][]int, [][]int) {
	rules := make(map[string][]int)
	var scannnedTickets [][]int

	// scan over rules
	for _, line := range input {
		if line == "" {
			break
		}
		kv := strings.Split(line, ": ")
		field := kv[0]
		// extract ranges
		var ranges []int
		rangeNotes := strings.Split(kv[1], " or ")
		for _, r := range rangeNotes {
			minmax := strings.Split(r, "-")
			for _, a := range minmax {
				i, err := strconv.Atoi(a)
				if err != nil {
					log.Fatalf("Failed to parse line %q, %#v", line, err)
				}
				ranges = append(ranges, i)
			}
		}
		rules[field] = ranges
	}

	// scan over tickets
	nearbySection := false
	for _, line := range input {
		if strings.Contains(line, "nearby tickets:") {
			nearbySection = true
			continue
		}
		if !nearbySection {
			continue
		}

		var ticketFields []int
		split := strings.Split(line, ",")
		for _, a := range split {
			i, err := strconv.Atoi(a)
			if err != nil {
				log.Fatalf("Failed to parse line %q, %#v", line, err)
			}
			ticketFields = append(ticketFields, i)
		}
		scannnedTickets = append(scannnedTickets, ticketFields)
	}
	return rules, scannnedTickets
}
