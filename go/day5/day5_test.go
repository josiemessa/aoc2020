package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"BFFFBBFRRR", //row 70, column 7, seat ID 567.
	"FFFBBBFRRR", // row 14, column 7, seat ID 119.
	"BBFFBBFRLL", // row 102, column 4, seat ID 820.
}

func TestDay5Part1(t *testing.T) {
	require.Equal(t, 820.0, SolveP1(input))
}

func Test_partitionToSeat(t *testing.T) {
	type args struct {
		par string
	}
	tests := []struct {
		name       string
		args       args
		wantRow    float64
		wantColumn float64
	}{
		{name: "line 1", args: args{"BFFFBBFRRR"},
			wantRow: 70, wantColumn: 7},
		{name: "line 2", args: args{"FFFBBBFRRR"},
			wantRow: 14, wantColumn: 7},
		{name: "line 3", args: args{"BBFFBBFRLL"},
			wantRow: 102, wantColumn: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRow, gotColumn := partitionToSeat(tt.args.par)
			require.Equal(t, tt.wantRow, gotRow)
			require.Equal(t, tt.wantColumn, gotColumn)
		})
	}
}
