using System;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Day4
{
    public class PassportParser
    {
        private string _input;

        public PassportParser(string input)
        {
            _input = input;
        }

        // public int Solve()
        // {
        //     // sort into individual passport definitions
        //     var lines = _input.Split(Environment.NewLine);
        //     var passports = new List<string>();
        //     var currentPassport = "";
        //     foreach (var line in lines)
        //     {
        //         if (line.Length != 0)
        //         {
        //             currentPassport += $" {line}";
        //             continue;
        //         }
        //
        //         passports.Add(currentPassport);
        //         currentPassport = "";
        //     }
        //
        //     return passports.Where(s => s.Count(c => c == ':') == 8)
        // }

        // private static IEnumerable<Dictionary<string, string>> ParseData(string input, bool validate)
        // {
        //     return input.Split("\n\n")
        //         .Select(line => (line.Replace("\n", " ").Trim()).Split(" "))
        //         .Select(entries => entries.Select(v => v.Split(":"))
        //             .Where(pair => validate == false || ValidateEntry(pair[0], pair[1]))
        //             .ToDictionary(pair => pair[0], pair => pair[1]));
        // }
    }
}