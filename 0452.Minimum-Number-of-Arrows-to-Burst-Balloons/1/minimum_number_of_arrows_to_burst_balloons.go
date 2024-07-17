package minimum_number_of_arrows_to_burst_balloons

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 0
	maxThreshold := points[0][1]
	for _, point := range points {
		if maxThreshold < point[0] {
			maxThreshold = point[1]
			count++
		}
	}

	return count + 1
}
