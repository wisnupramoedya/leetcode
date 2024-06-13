package minimum_number_of_moves_to_seat_everyone

/*
Runtime: 0 ms
Memory Usage: 3.2 MB
*/

import (
	"math"
	"sort"
)

func deviate(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	sort.Slice(students, func(i, j int) bool {
		return students[i] < students[j]
	})

	n, moves := len(seats), 0

	for i := 0; i < n; i++ {
		moves += deviate(seats[i], students[i])
	}

	return moves
}
