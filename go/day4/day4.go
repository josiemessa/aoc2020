package day4

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Valid  bool
	fields int
	cid    bool
}

func SolveP1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	valid := 0
	totalFields := 0
	cid := false
	for scanner.Scan() {
		line := scanner.Text()

		if fields := strings.Count(line, ":"); fields != 0 {
			cid = cid || strings.Contains(line, "cid")
			totalFields += fields
			continue
		}

		// we've hit a new line so sum up what we've got so far
		if totalFields == 7 && !cid {
			valid++
		}
		if totalFields == 8 {
			valid++
		}

		// reset
		totalFields = 0
		cid = false
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error", err)
	}
	return valid
}

func SolveP2(input io.Reader) (valid int) {
	scanner := bufio.NewScanner(input)
	currentPassport := &Passport{Valid: true}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			// End of passport, validate
			if currentPassport.Valid && (currentPassport.fields == 8 || (currentPassport.fields == 7 && !currentPassport.cid)) {
				valid++
			}
			// reset
			currentPassport = &Passport{Valid: true}
		}
		for _, f := range fields {
			currentPassport.Valid = currentPassport.Valid && currentPassport.Parse(f) == nil
		}
	}
	// End of passport, validate
	if currentPassport.Valid && (currentPassport.fields == 8 || (currentPassport.fields == 7 && !currentPassport.cid)) {
		valid++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error", err)
	}
	return
}

var (
	hgtRegex = regexp.MustCompile(`(\d+)(cm)?(in)?`)
	hclRegex = regexp.MustCompile(`#[0-9a-f]{6}`)
)

func (p *Passport) Parse(field string) error {
	split := strings.Split(field, ":")
	if len(split) != 2 {
		log.Fatalf("failed to split field %q", field)
	}
	switch split[0] {
	case "byr":
		//byr (Birth Year) - four digits; at least 1920 and at most 2002.
		i, err := strconv.Atoi(split[1])
		if err != nil {
			return fmt.Errorf("failed to parse int value in field %q", field)
		}
		if !(i >= 1920 && i <= 2002) {
			return fmt.Errorf("birth year in field %q not valid %d", field, i)
		}
		p.fields++
	case "iyr":
		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		i, err := strconv.Atoi(split[1])
		if err != nil {
			return fmt.Errorf("failed to parse int value in field %q", field)
		}
		if !(i >= 2010 && i <= 2020) {
			return fmt.Errorf("issue year in field %q not valid", field)
		}
		p.fields++
	case "eyr":
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		i, err := strconv.Atoi(split[1])
		if err != nil {
			return fmt.Errorf("failed to parse int value in field %q", field)
		}
		if !(i >= 2020 && i <= 2030) {
			return fmt.Errorf("exp year in field %q not valid", field)
		}
		p.fields++
	case "hgt":
		// hgt (Height) - a number followed by either cm or in:
		//If cm, the number must be at least 150 and at most 193.
		//If in, the number must be at least 59 and at most 76.
		matches := hgtRegex.FindStringSubmatch(split[1])
		if len(matches) != 4 || matches[0] != split[1] {
			return fmt.Errorf("height in field %q not valid %#v", field, matches)
		}
		i, err := strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("failed to parse height measurement in field %q", field)
		}
		if !(matches[2] == "cm" && (150 <= i && i <= 193) || matches[3] == "in" && (59 <= i && i <= 76)) {
			return fmt.Errorf("height in field %q is invalid %#v", field, matches)
		}
		p.fields++
	case "hcl":
		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		if !hclRegex.MatchString(split[1]) {
			return fmt.Errorf("hair colour in field %q not valid", field)
		}
		p.fields++
	case "ecl":
		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		if len(split[1]) != 3 || !strings.Contains("amb blu brn gry grn hzl oth", split[1]) {
			return fmt.Errorf("eye colour in field %q not valid", field)
		}
		p.fields++
	case "pid":
		if _, err := strconv.Atoi(split[1]); err != nil || len(split[1]) != 9 {
			return fmt.Errorf("passport ID in field %q not valid", field)
		}
		p.fields++
	case "cid":
		p.cid = true
		p.fields++
	}
	return nil
}
