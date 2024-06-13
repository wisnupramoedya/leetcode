package kth_largest_element_in_an_array

/*
Runtime: 66 ms
Memory Usage: 11.4 MB
*/

func findKthLargest(nums []int, k int) int {
	// 1. find the range of nums
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		} else if max < num {
			max = num
		}
	}

	// 2. in freq array, count each show
	numRangeLen := max - min + 1
	numRanges := make([]int, numRangeLen)
	for _, num := range nums {
		numRanges[num-min]++
	}

	// 3. check the kth largest from freq array
	i := 0
	if k > len(nums)/2 {
		// if k > half of nums, start from the smallest
		for i, k = 0, len(nums)-k+1; i != numRangeLen; i++ {
			k -= numRanges[i]
			if k <= 0 {
				break
			}
		}
	} else {
		// if k <= half of nums, start from the largest
		for i = numRangeLen - 1; i != -1; i-- {
			k -= numRanges[i]
			if k <= 0 {
				break
			}
		}
	}

	return i + min
}
