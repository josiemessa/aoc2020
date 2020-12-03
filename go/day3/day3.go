package day3

type TobogganSolver struct {
	Input []string
	Slope SlopeDefinition
}
type SlopeDefinition map[Direction]int

type Direction int

const (
	Right Direction = iota
	Down
)

// Solve returns the number of trees encountered using the route
func (s *TobogganSolver) Solve() int {
	if s.Slope[Right] == 0 || s.Slope[Down] == 0 {
		panic("josie pls")
	}

	// precompute location of trees
	treeMap := make(map[vector2]struct{})
	s.lookForTrees("#"[0], treeMap)

	// we start at the top left, so assume that the first element in the first line of input is (0,0),
	// the second element in the first line of input is (1, 0)
	// the first element in the second line of input is (0, 1), and
	// the second element in the second line of input is (1,1)
	// also can't figure out why th
	pos := vector2{x: 0, y: 0}
	trees := 0
	for pos.y < len(s.Input) {
		// Check current position first
		if _, ok := treeMap[pos]; ok {
			trees++
		}
		// Move position so for loop check runs immediately after
		pos = s.move(pos)
	}
	return trees
}

type vector2 struct {
	x int
	y int
}

// for some reason I can't return the map here and I cba to figure out why so we're C now
func (s *TobogganSolver) lookForTrees(tree uint8, output map[vector2]struct{}) {
	for i, line := range s.Input {
		for j, char := range line {
			if uint8(char) == tree {
				// i is the row, so the position on the y axis, j is the column, position on x
				output[vector2{j, i}] = struct{}{}
			}
		}
	}
}

func (s *TobogganSolver) move(pos vector2) vector2 {
	// Moving to the right increases x, moving down increases y
	newPos := vector2{pos.x + s.Slope[Right], pos.y + s.Slope[Down]}

	// Input repeats infinitely to the right, so mod x position against length of line (they are uniform so we can do this)
	newPos.x = newPos.x % (len(s.Input[0]))
	return newPos
}
