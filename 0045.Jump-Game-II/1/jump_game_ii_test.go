package jump_game_ii

import "testing"

func Test45_1_1(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	output := 2

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test45_1_2(t *testing.T) {
	nums := []int{2, 3, 0, 1, 4}
	output := 2

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test45_1_3(t *testing.T) {
	nums := []int{2, 2, 4, 1, 4}
	output := 2

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test45_1_4(t *testing.T) {
	nums := []int{2, 3, 1}
	output := 1

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test45_1_5(t *testing.T) {
	nums := []int{3, 2, 1}
	output := 1

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test45_1_6(t *testing.T) {
	nums := []int{1, 2, 1, 1, 1}
	output := 3

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test45_1_7(t *testing.T) {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}
	output := 2

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test45_1_8(t *testing.T) {
	nums := []int{4, 1, 1, 3, 1, 1, 1}
	output := 2

	result := jump(nums)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}
