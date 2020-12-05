package day5

import (
	"math"
	"sort"
)

func SolveP1(input []string) float64 {
	var maxID float64
	for _, seat := range input {
		row, column := partitionToSeat(seat)
		id := seatToID(row, column)
		maxID = math.Max(id, maxID)
	}
	return maxID
}

func SolveP2(input []string) float64 {
	var IDs []float64

	for _, ticket := range input {
		row, column := partitionToSeat(ticket)
		IDs = append(IDs, seatToID(row, column))
		sort.Float64s(IDs)
	}
	for i, ID := range IDs {
		if i < 2 {
			continue
		}
		if ID-IDs[i-1] == 2 {
			return ID - 1
		}
	}
	return 0.0
}

// BFFFBBFRRR: row 70, column 7, seat ID 567.
// FFFBBBFRRR: row 14, column 7, seat ID 119.
// BBFFBBFRLL: row 102, column 4, seat ID 820.
func partitionToSeat(par string) (row float64, column float64) {
	for i, char := range par {
		if char == 'F' || char == 'L' {
			continue
		}
		if char == 'B' {
			row += math.Exp2(float64(6 - i))
			continue
		}
		// assume char == 'R' and everything is well formed
		// we take 9 off this  as we're in indices 7,8 and 9
		column += math.Exp2(float64(9 - i))
	}
	return
}

func seatToID(row float64, column float64) float64 {
	return row*8 + column
}
