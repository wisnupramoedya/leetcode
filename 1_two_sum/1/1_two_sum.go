package __two_sum

/*
	Runtime: 6 ms, faster than 70.44% of Go online submissions for Two Sum.
	Memory Usage: 4.2 MB, less than 57.25% of Go online submissions for Two Sum.
*/

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	var first, second int

	for index, num := range nums {
		lastIndex, exists := set[num]

		if exists {
			first, second = lastIndex, index
			break
		} else {
			set[target-num] = index
		}

	}
	return []int{first, second}
}
