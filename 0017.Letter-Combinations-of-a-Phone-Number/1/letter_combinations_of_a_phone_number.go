package letter_combinations_of_a_phone_number

/*
Runtime: 0 ms
Memory Usage: 2.3 MB
*/

var buttons = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func combinePhone(array *[]string, digits *string, letters string) {
	index := len(letters)
	totalDigit := len(*digits)
	if totalDigit <= 0 {
		*array = []string{}
		return
	}

	for _, letter := range buttons[(*digits)[index]] {
		if len(letters) == totalDigit-1 {
			*array = append(*array, letters+string(letter))
			continue
		}

		combinePhone(array, digits, letters+string(letter))
	}
}

func letterCombinations(digits string) []string {
	var output []string

	combinePhone(&output, &digits, "")
	return output
}
