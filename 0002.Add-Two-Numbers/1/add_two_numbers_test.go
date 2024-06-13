package __add_two_numbers

import (
	"github.com/wisnupramoedya/leetcode/structures"
	"github.com/wisnupramoedya/leetcode/utils"
	"testing"
)

func run(l1, l2 []int) []int {
	l1ListNode, l2ListNode := structures.SetArrayInt2ListNode(l1), structures.SetArrayInt2ListNode(l2)
	resultListNode := addTwoNumbers(l1ListNode, l2ListNode)
	return structures.SetListNode2ArrayInt(resultListNode)
}

func Test2_1_1(t *testing.T) {
	l1, l2 := []int{2, 4, 3}, []int{5, 6, 4}
	output := []int{7, 0, 8}

	result := run(l1, l2)

	if !utils.IsArrayNumberSame(output, result) {
		t.Errorf("It is not the same array of int")
	}
}

func Test2_1_2(t *testing.T) {
	l1, l2 := []int{0}, []int{0}
	output := []int{0}

	result := run(l1, l2)

	if !utils.IsArrayNumberSame(output, result) {
		t.Errorf("It is not the same array of int")
	}
}

func Test2_1_3(t *testing.T) {
	l1, l2 := []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}
	output := []int{8, 9, 9, 9, 0, 0, 0, 1}

	result := run(l1, l2)

	if !utils.IsArrayNumberSame(output, result) {
		t.Errorf("It is not the same array of int")
	}
}
