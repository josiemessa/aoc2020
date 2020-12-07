package day7

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// for each colour, I want to know the allowed contents as well as any containers
type bagNode struct {
	// contents is the bags allowed to exist inside this bag
	contents []bagEdge
	// containers is the colour of bags that can contain this bag
	containers []string
}

type bagEdge struct {
	colour   string
	quantity int
}

type luggageSorter struct {
	result map[string]struct{}
	graph  map[string]*bagNode
	rules  map[string][]bagEdge
	count  int
}

func Solve(colourMatch string, rulesDef string, part2 bool) int {
	s := &luggageSorter{}
	s.rulesDefToRules(rulesDef)
	s.rulesToGraph()

	if !part2 {
		s.search(colourMatch)
		return len(s.result)
	}

	// Take one off as we don't include the current bag
	return s.countContents(bagEdge{"shiny gold", 1}) - 1
}

func (s *luggageSorter) countContents(edge bagEdge) int {
	node := s.graph[edge.colour]

	// we need to calculate the running total at leach level of recursion

	// runningTotal is the sum of the weighted children +1 as we have to include the current bag
	runningTotal := 1
	// keep traversing
	for _, child := range node.contents {
		runningTotal += s.countContents(child)
	}
	return runningTotal * edge.quantity
}

func (s *luggageSorter) search(colourMatch string) {
	if s.result == nil {
		s.result = make(map[string]struct{})
	}

	containers := s.graph[colourMatch].containers

	for _, colour := range containers {
		s.result[colour] = struct{}{}
		s.search(colour)
	}
}

// A graph shows, for a given colour (the key to the graph), which
// coloured bags can be contained inside it (contents), and which coloured bags can contain that colour (containers)
func (s *luggageSorter) rulesToGraph() {
	s.graph = make(map[string]*bagNode)
	for parentColour, allowedContents := range s.rules {
		node, ok := s.graph[parentColour]
		if !ok {
			node = &bagNode{}
			s.graph[parentColour] = node
		}
		node.contents = allowedContents
		// Go through allowed contents and set containers
		for _, bag := range allowedContents {
			// look up the node for the current colour's bag and set containers
			thisBagNode, ok := s.graph[bag.colour]
			if !ok {
				// we haven't found this colour yet so start building it up
				s.graph[bag.colour] = &bagNode{
					containers: []string{parentColour},
				}
			} else {
				thisBagNode.containers = append(thisBagNode.containers, parentColour)
			}
		}
	}
}

var bagRegex = regexp.MustCompile(`(\d+?) (\w+? \w+?) bag`)

func (s *luggageSorter) rulesDefToRules(rulesDef string) {
	s.rules = make(map[string][]bagEdge)
	for _, line := range strings.Split(rulesDef, "\n") {
		kv := strings.Split(line, " bags contain ")
		s.rules[kv[0]] = []bagEdge{}
		if strings.TrimSpace(kv[1]) == "no other bags." {
			continue
		}
		for _, bag := range strings.Split(kv[1], ",") {
			matches := bagRegex.FindStringSubmatch(bag)
			if len(matches) != 3 {
				log.Fatalf("Couldn't match line %v, split %s, found %#v", kv[1], bag, matches)
			}
			i, _ := strconv.Atoi(matches[1]) // ignore error as it wouldn't have matched if it wasn't an int
			s.rules[kv[0]] = append(s.rules[kv[0]],
				bagEdge{
					colour:   matches[2],
					quantity: i})
		}
	}
}
