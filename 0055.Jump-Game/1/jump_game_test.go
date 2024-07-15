package jump_game

import "testing"

func Test55_1_1(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	output := true

	result := canJump(nums)

	if output != result {
		t.Errorf("Expected %t, but got %t", output, result)
	}
}

func Test55_1_2(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	output := false

	result := canJump(nums)

	if output != result {
		t.Errorf("Expected %t, but got %t", output, result)
	}
}

func Test55_1_3(t *testing.T) {
	nums := []int{0}
	output := true

	result := canJump(nums)

	if output != result {
		t.Errorf("Expected %t, but got %t", output, result)
	}
}
