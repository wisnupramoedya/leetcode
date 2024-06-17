package letter_combinations_of_a_phone_number

/*
Runtime: 0 ms
Memory Usage: 2.4 MB
*/

var buttons = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var combinations []string
	for i, digit := range digits {
		// for first iteration
		if i == 0 {
			combinations = buttons[digit]
			continue
		}

		// if next digit iteration
		var newCombinations []string
		for _, combination := range combinations {
			for _, key := range buttons[digit] {
				newCombinations = append(newCombinations, combination+key)
			}
		}
		combinations = newCombinations
	}
	return combinations
}
