using System;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode2020
{
    public class ExpenseSolver
    {
        public List<int> Inputs { get; set; }


        public int Solve2(int target)
        {
            var pair = Inputs.Where(a => Inputs.Any(b => a + b == target && a != b));

            var result = 1;
            foreach (var y in pair)
            {
                Console.WriteLine(y);
                result *= y;
            }

            return result;
        }

        public int Solve3(int target)
        {
            var triple = Inputs.Where(a => Inputs.Any(b =>
                Inputs.Any(c => a + b + c == target && a != b && a != c)));

            var result = 1;
            foreach (var y in triple)
            {
                Console.WriteLine(y);
                result *= y;
            }

            return result;
        }
    }
}