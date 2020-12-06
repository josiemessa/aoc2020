using AdventOfCode.Day6;
using Xunit;

namespace AdventOfCode2020.Tests
{
    public class CustomsCheckTest
    {
            string input = @"abc

a
b
c

ab
ac

a
a
a
a

b";
        [Fact]
        public void Solve_WithExampleInput()
        {

            var check = new CustomsCheck(input);
            Assert.Equal(11, check.Solve(false));
        }

        [Fact]
        public void Solve_WithAndCheck_WithExampleInput()
        {
            var check = new CustomsCheck(input);
            Assert.Equal(6, check.Solve(true));
        }
    }
}