using System;
using System.Collections.Generic;
using System.Data;
using System.Linq;

namespace AdventOfCode.Day6
{
    public class CustomsCheck
    {
        private string _input;
        private int _solution;

        public CustomsCheck(string input)
        {
            _input = input;
        }

        public int Solve(bool andCheck)
        {
            _solution = 0;
            var groupAnswers = _input.Split($"{Environment.NewLine}{Environment.NewLine}") // split into groups of answers
                //within a group, put each person's answers into its own element in an array
                .Select(line => (line.Replace($"{Environment.NewLine}", " ").Trim()).Split(" "));

            foreach (var group in groupAnswers)
            {
                if (andCheck)
                {
                    GroupAndCheck(@group);
                }
                else
                {
                    GroupOrCheck(group);
                }
            }

            return _solution;
        }

        private void GroupOrCheck(IEnumerable<string> group)
        {
            var yesses = new HashSet<char>();
            foreach (var person in group)
            {
                foreach (var answer in person)
                {
                    yesses.Add(answer);
                }
            }

            _solution += yesses.Count;
        }

        private void GroupAndCheck(string[] @group)
        {
            var foundAnswers = @group[0].ToList();
            foreach (var person in group)
            {
                var copy = foundAnswers.Select(item => (char)item).ToList();
                foreach (var answer in foundAnswers.Where(answer => !person.Contains(answer)))
                {
                    copy.Remove(answer);
                }

                foundAnswers = copy;
            }

            _solution += foundAnswers.Count;
        }
    }
}