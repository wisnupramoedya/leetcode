package combination_sum

/*
Runtime: 3 ms
Memory Usage: 3.3 MB
*/

func combination(output *[][]int, candidates *[]int, combinations []int, position, target int) {
	if target == 0 {
		newCombinations := make([]int, len(combinations))
		copy(newCombinations, combinations)
		*output = append(*output, newCombinations)
		return
	}

	for i := position; i < len(*candidates); i++ {
		if target-(*candidates)[i] >= 0 {
			combination(output, candidates, append(combinations, (*candidates)[i]), i, target-(*candidates)[i])
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var output [][]int

	combination(&output, &candidates, []int{}, 0, target)

	return output
}
