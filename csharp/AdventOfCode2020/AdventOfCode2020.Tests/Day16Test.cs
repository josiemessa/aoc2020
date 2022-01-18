using System;
using Xunit;
using Xunit.Abstractions;

namespace AdventOfCode2020.Tests
{
    public class Day16Test
    {
        private readonly ITestOutputHelper _testOutputHelper;

        public Day16Test(ITestOutputHelper testOutputHelper)
        {
            _testOutputHelper = testOutputHelper;
        }

        [Fact]
        public void Rule_ValidateTest()
        {
            var input = new int[] {7, 3};
            var rule = new Rule("class", new int[2] {1, 5}, new int[] {3, 7});
            foreach (var i in input)
            {
                Assert.True(rule.Validate(i));
            }
            Assert.False(rule.Validate(70));

        }

        [Fact]
        public void Day16_SolveTest()
        {
            var input = @"class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9";
            var solver = new Day16(input);
            solver.Solve();
        }
    }
}