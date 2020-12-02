package day2

import (
	"log"
	"strconv"
	"strings"
)

type PasswordSolver struct {
	Input []string
}

type Policy interface {
	apply(string) bool
}

// Solve returns the number of valid passwords supplied in the list Input
func (p *PasswordSolver) Solve(old bool) int {
	validPasswords := 0
	for _, line := range p.Input {
		var (
			policy   Policy
			password string
		)
		policy, password = parseLine(line, old)
		valid := policy.apply(password)
		if valid {
			validPasswords++
		}
	}
	return validPasswords
}

type newPolicy struct {
	pos1      int
	pos2      int
	character string
}

func (p *newPolicy) apply(password string) bool {
	if len(password) < p.pos1 {
		return false
	}
	if string(password[p.pos1-1]) == p.character {
		// can't have more than one occurrence of character in string
		if len(password) >= p.pos2 && string(password[p.pos2-1]) == p.character {
			return false
		}
		return true
	}

	if len(password) >= p.pos2 && string(password[p.pos2-1]) == p.character {
		return true
	}

	return false
}

type oldPolicy struct {
	frequencyMin int
	frequencyMax int
	character    string
}

// Apply the oldPolicy to the provided password and return whether the password is valid or not
func (p *oldPolicy) apply(password string) bool {
	if len(password) < p.frequencyMin {
		return false
	}

	count := strings.Count(password, p.character)
	if count < p.frequencyMin || count > p.frequencyMax {
		return false
	}

	return true
}

func parseLine(line string, old bool) (Policy, string) {
	//"1-3 a: abcde"
	split := strings.Split(line, " ")
	if len(split) != 3 {
		log.Fatalf("expected three parts to line - line was: %q", line)
	}

	// Calculate min/max values. These are separated by a "-" so split on there and parse into int
	minmax := strings.Split(split[0], "-")
	if len(minmax) != 2 {
		log.Fatalf("expected two parts to min/max - string was: %q", split[0])
	}
	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		log.Fatalf("min was not an int - min was %q, err was %q", minmax[0], err)
	}

	max, err2 := strconv.Atoi(minmax[1])
	if err2 != nil {
		log.Fatalf("max was not an int - max was %q, err was %q", minmax[1], err2)
	}
	if old {

		policy := &oldPolicy{}
		policy.frequencyMin = min
		policy.frequencyMax = max

		// Calculate character - this is suffixed with a ":" so strip that out.
		character := strings.TrimSuffix(split[1], ":")
		if len(character) != 1 {
			log.Fatalf("character was not a single character - character was %q", split[1])
		}
		policy.character = character
		// Finally we assume the last string is the password, so just return that
		return policy, split[2]
	}
	policy := &newPolicy{}
	policy.pos1 = min
	policy.pos2 = max

	// Calculate character - this is suffixed with a ":" so strip that out.
	character := strings.TrimSuffix(split[1], ":")
	if len(character) != 1 {
		log.Fatalf("character was not a single character - character was %q", split[1])
	}
	policy.character = character
	// Finally we assume the last string is the password, so just return that
	return policy, split[2]
}
