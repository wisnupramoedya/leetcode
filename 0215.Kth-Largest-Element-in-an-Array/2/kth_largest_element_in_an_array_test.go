package kth_largest_element_in_an_array

import (
	"testing"
)

func Test215_1_1(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	expected := 5

	result := findKthLargest(nums, k)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func Test215_1_2(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4

	expected := 4

	result := findKthLargest(nums, k)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func Test215_1_3(t *testing.T) {
	nums := []int{2, 1}
	k := 2

	expected := 1

	result := findKthLargest(nums, k)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
