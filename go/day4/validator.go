package day4

import (
	"regexp"
	"strconv"
	"strings"
)

type validator func(string) bool

var (
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	//hgt (Height) - a number followed by either cm or in:
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	//pid (Passport ID) - a nine-digit number, including leading zeroes.
	//cid (Country ID) - ignored, missing or not.
	validators = map[string]validator{
		"byr": func(val string) bool { return numValidator(val, 1920, 2002) },
		"iyr": func(val string) bool { return numValidator(val, 2010, 2020) },
		"eyr": func(val string) bool { return numValidator(val, 2020, 2030) },
		"hgt": hgtValidator,
		"hcl": func(val string) bool { return hclRegex.MatchString(val) },
		"ecl": func(val string) bool { return len(val) == 3 && strings.Contains("amb blu brn gry grn hzl oth", val) },
		"pid": func(val string) bool { _, err := strconv.Atoi(val); return err == nil && len(val) == 9 },
		"cid": func(s string) bool { return true },
	}
	hgtRegex = regexp.MustCompile(`(\d+)(cm|in)`)
	hclRegex = regexp.MustCompile(`#[0-9a-f]{6}`)
)

func numValidator(val string, min, max int) bool {
	i, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return min <= i && i <= max
}

func hgtValidator(val string) bool {
	matches := hgtRegex.FindStringSubmatch(val)
	if len(matches) != 3 || matches[0] != val {
		return false
	}
	i, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}
	return (matches[2] == "cm" && (150 <= i && i <= 193)) || (matches[2] == "in" && (59 <= i && i <= 76))
}
