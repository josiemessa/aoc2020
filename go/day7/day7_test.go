package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var rules = map[string][]bagEdge{
	"light red": {
		bagEdge{colour: "bright white", quantity: 1},
		bagEdge{quantity: 2, colour: "muted yellow"},
	},
	"dark orange": {
		bagEdge{colour: "bright white", quantity: 3},
		bagEdge{quantity: 4, colour: "muted yellow"},
	},
	"bright white": {
		bagEdge{colour: "shiny gold", quantity: 1},
	},
	"muted yellow": {
		bagEdge{colour: "shiny gold", quantity: 2},
		bagEdge{quantity: 9, colour: "faded blue"},
	},
	"shiny gold": {
		bagEdge{colour: "dark olive", quantity: 1},
		bagEdge{quantity: 2, colour: "vibrant plum"},
	},
	"dark olive": {
		bagEdge{colour: "faded blue", quantity: 3},
		bagEdge{quantity: 4, colour: "dotted black"},
	},
	"vibrant plum": {
		bagEdge{colour: "faded blue", quantity: 5},
		bagEdge{quantity: 6, colour: "dotted black"},
	},
	"faded blue":   {},
	"dotted black": {},
}

func TestSolve(t *testing.T) {
	require.Equal(t, 4, Solve("shiny gold", ""))
}

func TestLuggageSorter_rulesToGraph(t *testing.T) {
	s := luggageSorter{}
	s.rules = rules
	s.rulesToGraph()
	require.Len(t, s.graph, len(s.rules))
}

func TestLuggageSorter_Search(t *testing.T) {
	s := luggageSorter{}
	s.rules = rules
	s.rulesToGraph()
	require.Len(t, s.graph, len(s.rules))
	s.search("shiny gold")
	require.Len(t, s.result, 4)
}

func TestLuggageSorter_RulesDefToRules(t *testing.T) {
	s := luggageSorter{}
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	s.rulesDefToRules(input)
	require.EqualValues(t, rules, s.rules)
}
