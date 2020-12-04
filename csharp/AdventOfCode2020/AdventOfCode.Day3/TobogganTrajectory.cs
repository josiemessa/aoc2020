using System;
using System.Collections.Generic;

namespace AdventOfCode2020
{
    public class TobogganTrajectory
    {
        private List<string> _input;

        public TobogganTrajectory(List<string> input)
        {
            _input = input;
        }

        public int Solve(Dictionary<Direction, int> slope)
        {
            if (slope[Direction.Right] == 0 || slope[Direction.Down] == 0)
            {
                throw new Exception("Slope invalid");
            }

            foreach (var line in _input)
            {

            }
        }

    }
    public enum Direction
    {
        Right,
        Down
    }


}