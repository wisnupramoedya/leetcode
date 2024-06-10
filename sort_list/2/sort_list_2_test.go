package sort_list2

import (
	"github.com/wisnupramoedya/leetcode"
	"testing"
)

func Test148_2_1(t *testing.T) {
	head := []int{4, 2, 1, 3}
	expected := []int{1, 2, 3, 4}

	headLinkedList := setLinkedList(head)

	result := sortList(headLinkedList)
	resultInt := setLinkedListToArrayInt(result)

	if !leetcode.IsArrayIntSame(resultInt, expected) {
		t.Errorf("It is not the same array of int")
	}
}

func Test148_2_2(t *testing.T) {
	head := []int{-1, 5, 3, 4, 0}
	expected := []int{-1, 0, 3, 4, 5}

	headLinkedList := setLinkedList(head)

	result := sortList(headLinkedList)
	resultInt := setLinkedListToArrayInt(result)

	if !leetcode.IsArrayIntSame(resultInt, expected) {
		t.Errorf("It is not the same array of int")
	}
}

func Test148_2_3(t *testing.T) {
	var head []int
	var expected []int

	headLinkedList := setLinkedList(head)

	result := sortList(headLinkedList)
	resultInt := setLinkedListToArrayInt(result)

	if !leetcode.IsArrayIntSame(resultInt, expected) {
		t.Errorf("It is not the same array of int")
	}
}

func Test148_2_4(t *testing.T) {
	head := []int{-1}
	expected := []int{-1}

	headLinkedList := setLinkedList(head)

	result := sortList(headLinkedList)
	resultInt := setLinkedListToArrayInt(result)

	if !leetcode.IsArrayIntSame(resultInt, expected) {
		t.Errorf("It is not the same array of int")
	}
}
