package sort_list1

import (
	"github.com/wisnupramoedya/leetcode"
	"testing"
)

func Test148_1_1(t *testing.T) {
	head := []int{4, 2, 1, 3}
	expected := []int{1, 2, 3, 4}

	headLinkedList := SetLinkedList(head)

	result := SortList(headLinkedList)
	resultInt := SetLinkedListToArrayInt(result)

	if !leetcode.IsArrayIntSame(resultInt, expected) {
		t.Errorf("It is not the same array of int")
	}
}

func Test148_1_2(t *testing.T) {
	head := []int{-1, 5, 3, 4, 0}
	expected := []int{-1, 0, 3, 4, 5}

	headLinkedList := SetLinkedList(head)

	result := SortList(headLinkedList)
	resultInt := SetLinkedListToArrayInt(result)

	if !leetcode.IsArrayIntSame(resultInt, expected) {
		t.Errorf("It is not the same array of int")
	}
}

func Test148_1_3(t *testing.T) {
	var head []int
	var expected []int

	headLinkedList := SetLinkedList(head)

	result := SortList(headLinkedList)
	resultInt := SetLinkedListToArrayInt(result)
	println(resultInt)

	if !leetcode.IsArrayIntSame(resultInt, expected) {
		t.Errorf("It is not the same array of int")
	}
}
