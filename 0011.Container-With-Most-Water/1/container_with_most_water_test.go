package container_with_most_water

import "testing"

func Test11_1_1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := 49

	result := maxArea(height)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test11_1_2(t *testing.T) {
	height := []int{1, 1}
	output := 1

	result := maxArea(height)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test11_1_3(t *testing.T) {
	height := []int{1, 4, 5, 5, 4, 1}
	output := 12

	result := maxArea(height)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}
