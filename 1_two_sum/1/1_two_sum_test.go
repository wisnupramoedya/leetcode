package __two_sum

import "testing"

func Test1_1_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	output := []int{0, 1}

	result := twoSum(nums, target)

	for i := 0; i < 2; i++ {
		if result[i] != output[i] {
			t.Errorf("Expected %d, but got %d in index %d", output[i], result[i], i)
			break
		}
	}
}

func Test1_1_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	output := []int{1, 2}

	result := twoSum(nums, target)

	for i := 0; i < 2; i++ {
		if result[i] != output[i] {
			t.Errorf("Expected %d, but got %d in index %d", output[i], result[i], i)
			break
		}
	}
}

func Test1_1_3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	output := []int{0, 1}

	result := twoSum(nums, target)

	for i := 0; i < 2; i++ {
		if result[i] != output[i] {
			t.Errorf("Expected %d, but got %d in index %d", output[i], result[i], i)
			break
		}
	}
}
