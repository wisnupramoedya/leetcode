package combination_sum

import (
	"github.com/wisnupramoedya/leetcode/utils"
	"testing"
)

func Test39_1_1(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7

	output := [][]int{{2, 2, 3}, {7}}

	result := combinationSum(candidates, target)

	if !utils.Is2DArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}

func Test39_1_2(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8

	output := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}

	result := combinationSum(candidates, target)

	if !utils.Is2DArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}

func Test39_1_3(t *testing.T) {
	candidates := []int{2}
	target := 1

	var output [][]int

	result := combinationSum(candidates, target)

	if !utils.Is2DArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}

func Test39_1_4(t *testing.T) {
	candidates := []int{7, 3, 2}
	target := 18

	output := [][]int{
		{7, 7, 2, 2},
		{7, 3, 3, 3, 2},
		{7, 3, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 2, 2, 2},
		{3, 3, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
	}

	result := combinationSum(candidates, target)

	if !utils.Is2DArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}
