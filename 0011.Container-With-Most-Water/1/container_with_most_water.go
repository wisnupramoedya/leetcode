package container_with_most_water

/*
Runtime: 57 ms
Memory Usage: 7.7 MB
*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	areaTotal := 0

	for {
		if left == right {
			break
		}

		if height[left] < height[right] {
			areaTotal = max(areaTotal, height[left]*(right-left))
			left++
		} else {
			areaTotal = max(areaTotal, height[right]*(right-left))
			right--
		}
	}
	return areaTotal
}
