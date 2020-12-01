using System;
using System.Collections.Generic;
using Xunit;

namespace AdventOfCode2020.Tests
{
    public class ExpenseSolverTest
    {
        [Fact]
        public void Solve2_WithExampleInput_ShouldSolve()
        {
            var solver = new ExpenseSolver
            {
                Inputs = new List<int> {1721, 979, 366, 299, 675, 1456}
            };
            Assert.Equal(514579, solver.Solve2(2020));
        }

        [Fact]
        public void Solve3_WithExampleInput_ShouldSolve()
        {
            var solver = new ExpenseSolver
            {
                Inputs = new List<int> {1721, 979, 366, 299, 675, 1456}
            };
            Assert.Equal(241861950, solver.Solve3(2020));
        }
    }
}