using System;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode2020
{
    public class Day16
    {
        private List<Rule> _rules;
        private List<List<int>> _nearbyTickets;
        private IEnumerable<int> _yourTicket;
        private List<List<int>> _validTickets;

        public Day16(string input)
        {
            _rules = new List<Rule>();
            _validTickets = new List<List<int>>();
            ParseInput(input);
        }

        private void ParseInput(string input)
        {
            var sections = input.Split($"{Environment.NewLine}{Environment.NewLine}");

            // First section is rules
            ParseRules(sections[0]);

            // Second section is your ticket
            _yourTicket = sections[1].TrimStart($"your ticket:{Environment.NewLine}".ToCharArray()).Split(",").ToList()
                .Select(s => int.Parse(s));

            // Third section is nearby tickets
            _nearbyTickets = sections[2].TrimStart($"nearby tickets:{Environment.NewLine}".ToCharArray())
                .Split($"{Environment.NewLine}").ToList()
                .Select(line => line.Split(",").Select(int.Parse).ToList()).ToList();
        }

        public void Solve()
        {
            Console.WriteLine($"Part 1: {Part1()}");
            Console.WriteLine($"Found {_validTickets.Count} valid tickets");

            var finalPositions = FindRulePositions();
            foreach (var rule in finalPositions)
            {
                Console.WriteLine(rule.ToString());
            }
        }

        private Dictionary<int, Rule> FindRulePositions()
        {
            // organise ticket values by their index rather than the order they were in the input
            // so ticketsByIndex[0] is an array of the first value in a ticket.
            int[][] ticketsByIndex = new int[_validTickets[0].Count][];
            for (int i = 0; i < _validTickets.Count; i++)
            {
                for (int j = 0; j < _validTickets[0].Count; j++)
                {
                    if (i == 0) ticketsByIndex[j] = new int[_validTickets.Count];
                    ticketsByIndex[j][i] = _validTickets[i][j];
                }
            }

            var finalPositions = new Dictionary<int, Rule>();
            var validRulesByPosition = new Dictionary<int, List<Rule>>();

            // Iterate all ticket values by their position and find all rules that are valid for that position
            for (var i = 0; i < ticketsByIndex.Length; i++)
            {
                var validRules = FindValidRules(new List<int>(ticketsByIndex[i]), finalPositions);

                // While we're here, check whether there is only one matching rule as we can now start to ignore that one
                if (validRules.Count == 1) finalPositions[i] = validRules[0];

                validRules = validRules.Where(rule => !finalPositions.ContainsValue(rule)).ToList();
                validRulesByPosition[i] = validRules;
            }

            Console.WriteLine($"found {finalPositions.Count} rules in first pass");

            // Now start iterating until we find all the final positions for each rule in the ticket.
            var remainingRules = _rules.Where(rule => !finalPositions.ContainsValue(rule)).ToList();
            while (remainingRules.Any())
            {
                // Only look at the remaining rules
                foreach (var rule in remainingRules)
                {
                    var index = FindUniqueRulePosition(rule, validRulesByPosition, finalPositions);
                    if (index > 0)
                    {
                        finalPositions[index] = rule;
                    }
                }
                // update which rules we check through after each pass
                remainingRules = _rules.Where(rule => !finalPositions.ContainsValue(rule)).ToList();
                Console.WriteLine($"There are {remainingRules.Count} remaining rules to find");
            }

            if (finalPositions.Count != _rules.Count) Console.WriteLine($"Only found {finalPositions.Count} positions");

            return finalPositions;
        }

        private int Part1()
        {
            var errorRate = new List<int>();
            foreach (var ticket in _nearbyTickets)
            {
                var errors = ticket.Where(i => _rules.All(rule => !rule.Validate(i))).ToList();
                errors.ForEach(i => errorRate.Add(i));
                if (errors.Count == 0) _validTickets.Add(ticket);
            }

            return errorRate.Sum();
        }

        private static int FindUniqueRulePosition(Rule rule, Dictionary<int, List<Rule>> validRulesByPosition, Dictionary<int,Rule> finalRules)
        {
            var count = 0;
            var index = -1;
            for (var i = 0; i < validRulesByPosition.Count; i++)
            {
                if (!validRulesByPosition[i].Contains(rule)) continue;
                if (validRulesByPosition[i].Count == 1) return index;
                count++;
                index = i;
            }

            foreach (var posRulePair in finalRules)
            {
                foreach (var posListPair in validRulesByPosition.Where(posListPair => posListPair.Value.Contains(posRulePair.Value)))
                {
                    posListPair.Value.Remove(posRulePair.Value);
                }
            }

            if (count == 1) return index;
            return -1;
        }

        private List<Rule> FindValidRules(List<int> ticketVals, Dictionary<int, Rule> finalPositions)
        {
            foreach (var rule in _rules)
            {
                var valid = true;
                foreach (var ticketVal in ticketVals)
                {
                    valid = valid && rule.Validate(ticketVal);
                }
            }

            var validRules = _rules.Where(rule => ticketVals.All(rule.Validate)).ToList();
            if (validRules.Count == 0)
            {
                throw new Exception($"No rules match position");
            }

            // remove any rules that we've already placed in their final positions
            validRules = validRules.Where(r => !finalPositions.ContainsValue(r)).ToList();

            return validRules;
        }

        private void ParseRules(string section)
        {
            var lines = section.Split("\n");
            foreach (var line in lines)
            {
                var kv = line.Split(": ");
                var ranges = kv[1].Split(" or ");
                var mins = new int[2];
                var maxes = new int[2];
                for (var index = 0; index < ranges.Length; index++)
                {
                    var range = ranges[index];
                    var minmax = range.Split("-");
                    mins[index] = int.Parse(minmax[0]);
                    maxes[index] = int.Parse(minmax[1]);
                }

                _rules.Add(new Rule(kv[0], mins, maxes));
            }
        }
    }

    public class Rule
    {
        private string _name;
        private int[] _mins;
        private int[] _maxes;

        public Rule(string name, int[] mins, int[] maxes)
        {
            _name = name;
            _mins = mins;
            _maxes = maxes;
        }

        public bool Validate(int num)
        {
            return (_mins[0] <= num && num <= _maxes[0]) || (_mins[1] <= num && num <= _maxes[1]);
        }

        public override string ToString()
        {
            return $"{_name}: {_mins[0]}-{_maxes[0]} or {_mins[1]}-{_maxes[1]}";
        }
    }
}