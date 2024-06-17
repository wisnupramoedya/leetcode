package letter_combinations_of_a_phone_number

import (
	"github.com/wisnupramoedya/leetcode/utils"
	"testing"
)

func Test17_1_1(t *testing.T) {
	digits := "23"
	output := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	result := letterCombinations(digits)

	if !utils.IsArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}

func Test17_1_2(t *testing.T) {
	digits := ""
	var output []string

	result := letterCombinations(digits)

	if !utils.IsArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}

func Test17_1_3(t *testing.T) {
	digits := "2"
	output := []string{"a", "b", "c"}

	result := letterCombinations(digits)

	if !utils.IsArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}
