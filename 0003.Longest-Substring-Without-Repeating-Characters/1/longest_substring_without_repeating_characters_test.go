package longest_substring_without_repeating_characters

import "testing"

func Test3_1_1(t *testing.T) {
	s := "abcabcbb"
	output := 3

	result := lengthOfLongestSubstring(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test3_1_2(t *testing.T) {
	s := "bbbbb"
	output := 1

	result := lengthOfLongestSubstring(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test3_1_3(t *testing.T) {
	s := "pwwkew"
	output := 3

	result := lengthOfLongestSubstring(s)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}
