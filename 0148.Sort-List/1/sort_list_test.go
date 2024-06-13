package sort_list

import (
	"github.com/wisnupramoedya/leetcode/structures"
	"github.com/wisnupramoedya/leetcode/utils"
	"testing"
)

func run(head []int) []int {
	headLinkedList := structures.SetArrayInt2ListNode(head)
	result := sortList(headLinkedList)

	return structures.SetListNode2ArrayInt(result)

}

func Test148_1_1(t *testing.T) {
	head := []int{4, 2, 1, 3}
	expected := []int{1, 2, 3, 4}

	result := run(head)

	if !utils.IsArrayNumberSame(result, expected) {
		t.Errorf("It is not the same array of int")
	}
}

func Test148_1_2(t *testing.T) {
	head := []int{-1, 5, 3, 4, 0}
	expected := []int{-1, 0, 3, 4, 5}

	result := run(head)

	if !utils.IsArrayNumberSame(result, expected) {
		t.Errorf("It is not the same array of int")
	}
}

func Test148_1_3(t *testing.T) {
	var head []int
	var expected []int

	result := run(head)

	if !utils.IsArrayNumberSame(result, expected) {
		t.Errorf("It is not the same array of int")
	}
}
