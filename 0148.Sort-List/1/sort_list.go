package sort_list

import "github.com/wisnupramoedya/leetcode/structures"

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func SetLinkedList(values []int) *ListNode {
	var firstListNode, prevListNode *ListNode
	for _, value := range values {
		//println("SL:", value)
		listNode := &ListNode{
			Val:  value,
			Next: nil,
		}
		if firstListNode == nil {
			firstListNode = listNode
		}
		if prevListNode != nil {
			prevListNode.Next = listNode
		}
		prevListNode = listNode
	}
	return firstListNode
}

func merge(itemsA, itemsB []int) []int {
	i, j := 0, 0
	var newItems []int
	for i < len(itemsA) && j < len(itemsB) {
		if itemsA[i] <= itemsB[j] {
			newItems = append(newItems, itemsA[i])
			i++
		} else {
			newItems = append(newItems, itemsB[j])
			j++
		}
	}
	for i < len(itemsA) {
		newItems = append(newItems, itemsA[i])
		i++
	}
	for j < len(itemsB) {
		newItems = append(newItems, itemsB[j])
		j++
	}
	return newItems
}

func mergeSort(items []int) []int {
	length := len(items)
	if length < 2 {
		return items
	}

	itemsA := mergeSort(items[:length/2])
	itemsB := mergeSort(items[length/2:])
	return merge(itemsA, itemsB)
}

func sortList(head *ListNode) *ListNode {
	var values []int
	for {
		if head == nil {
			//println("Break")
			break
		}
		//println(head.Val, head)
		values = append(values, head.Val)
		head = head.Next
	}

	//println(values)

	sorted := mergeSort(values)

	return SetLinkedList(sorted)
}
