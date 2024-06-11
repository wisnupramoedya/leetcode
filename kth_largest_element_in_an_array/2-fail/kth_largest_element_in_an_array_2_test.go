package kth_largest_element_in_an_array

import (
	"testing"
)

func Test215_2_1(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4

	expected := 4

	result := findKthLargest2(nums, k)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func Test215_2_2(t *testing.T) {
	nums := []int{2, 1}
	k := 2

	expected := 1

	result := findKthLargest2(nums, k)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
