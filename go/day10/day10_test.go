package day10

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	inputstr := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`
	require.Equal(t, 220, SolveP1(strings.Split(inputstr, "\n")))
}

func TestSolveP2(t *testing.T) {
	inputstr := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`
	require.Equal(t, 19208, SolveP2(strings.Split(inputstr, "\n")))
}

func TestDumbExample(t *testing.T) {
	input := []string{"1", "2", "3", "5", "6"}
	SolveP2(input)
}
