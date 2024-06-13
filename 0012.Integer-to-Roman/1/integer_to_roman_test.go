package integer_to_roman

import "testing"

func Test12_1_1(t *testing.T) {
	num := 3749
	output := "MMMDCCXLIX"

	result := intToRoman(num)

	if output != result {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}

func Test12_1_2(t *testing.T) {
	num := 58
	output := "LVIII"

	result := intToRoman(num)

	if output != result {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}

func Test12_1_3(t *testing.T) {
	num := 1994
	output := "MCMXCIV"

	result := intToRoman(num)

	if output != result {
		t.Errorf("Expected %s, but got %s", output, result)
	}
}
