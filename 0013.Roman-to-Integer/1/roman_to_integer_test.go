package roman_to_integer

import "testing"

func Test12_1_1(t *testing.T) {
	s := "III"
	output := 3

	result := romanToInt(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test12_1_2(t *testing.T) {
	s := "LVIII"
	output := 58

	result := romanToInt(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test12_1_3(t *testing.T) {
	s := "MCMXCIV"
	output := 1994

	result := romanToInt(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test12_1_4(t *testing.T) {
	s := "MMMDCCCXCIX"
	output := 3899

	result := romanToInt(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test12_1_5(t *testing.T) {
	s := "MMMCMXCIX"
	output := 3999

	result := romanToInt(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}
