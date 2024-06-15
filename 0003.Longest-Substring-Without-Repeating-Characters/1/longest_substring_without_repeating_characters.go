package longest_substring_without_repeating_characters

/*
Runtime: 0 ms
Memory Usage: 3.4 MB
*/

func lengthOfLongestSubstring(s string) int {
	set := map[int32]int{}
	left, maxLen, totalLen := 0, 0, 0

	for right, char := range s {
		//println("r:", right, "l:", left, "t:", totalLen, "m:", maxLen, "c:", string(char))
		if value, isExist := set[char]; isExist && value >= left {
			//println("is", isExist, "val:", value)
			left = value + 1
			totalLen = right - left + 1
			set[char] = right
			continue
		}

		set[char] = right
		totalLen++
		if maxLen < totalLen {
			maxLen = totalLen
		}
	}
	return maxLen
}
