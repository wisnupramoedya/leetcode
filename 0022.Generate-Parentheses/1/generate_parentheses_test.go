package generate_parentheses

import (
	"github.com/wisnupramoedya/leetcode/utils"
	"testing"
)

func Test22_1_1(t *testing.T) {
	n := 1
	output := []string{"()"}

	result := generateParenthesis(n)

	if !utils.IsArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}

func Test22_1_2(t *testing.T) {
	n := 2
	output := []string{"(())", "()()"}

	result := generateParenthesis(n)

	if !utils.IsArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}

func Test22_1_3(t *testing.T) {
	n := 3
	output := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	result := generateParenthesis(n)

	if !utils.IsArraySame(output, result) {
		t.Errorf("It is not the same array")
	}
}
