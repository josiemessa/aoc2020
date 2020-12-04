package day4

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type Passport struct {
	fields map[string]string
	valid  bool
}

func Solve(input io.Reader, extra bool) (valid int) {
	scanner := bufio.NewScanner(input)
	current := Passport{valid: true, fields: make(map[string]string)}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// new line, end of current passport. Clean up and reset
			if _, ok := current.fields["cid"]; len(current.fields) == 8 || len(current.fields) == 7 && !ok {
				valid++
			}
			current = Passport{valid: true, fields: make(map[string]string)}
			continue
		}

		// Part of passport definition, parse
		for _, f := range strings.Fields(line) {
			kv := strings.Split(f, ":")
			if !extra {
				current.fields[kv[0]] = kv[1]
				continue
			}
			valid, ok := validators[kv[0]]
			if !ok {
				continue
			}
			if valid(kv[1]) {
				current.fields[kv[0]] = kv[1]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan failed: %#v", err)
	}

	if _, ok := current.fields["cid"]; len(current.fields) == 8 || len(current.fields) == 7 && !ok {
		valid++
	}
	return
}
