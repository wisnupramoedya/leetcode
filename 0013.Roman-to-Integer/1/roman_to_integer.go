package roman_to_integer

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Roman to Integer.
Memory Usage: 3 MB, less than 92.44% of Go online submissions for Roman to Integer.
*/

func romanToInt(s string) int {
	result := 0
	romans := make(map[byte]int)
	romans['I'] = 1
	romans['V'] = 5
	romans['X'] = 10
	romans['L'] = 50
	romans['C'] = 100
	romans['D'] = 500
	romans['M'] = 1000

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romans[s[i]] < romans[s[i+1]] {
			result += romans[s[i+1]] - romans[s[i]]
			i++
			continue
		}

		result += romans[s[i]]
	}
	return result
}
