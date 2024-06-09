// tags: math
package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var firstNode, lastNode, listNode *ListNode
	var tens, i, sum int = 0, 0, 0
	var val1, val2 int = 0, 0
	for {
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
		}
		if l2 == nil {
			val2 = 0
		} else {
			val2 = l2.Val
		}
		sum = val1 + val2 + tens
		tens = int(sum / 10)
		listNode = &ListNode{sum % 10, nil}
		if lastNode != nil {
			lastNode.Next = listNode
		}
		lastNode = listNode
		if i == 0 {
			firstNode = listNode
		}

		if (l1 == nil && l2 == nil) || (!(l1 != nil && l1.Next != nil) && !(l2 != nil && l2.Next != nil)) {
			if tens != 0 {
				listNode = &ListNode{tens, nil}
				lastNode.Next = listNode
				lastNode = listNode
			}
			break
		} else {
			i++
			if l1 == nil || l1.Next == nil {
				l1 = nil
			} else {
				l1 = l1.Next
			}

			if l2 == nil || l2.Next == nil {
				l2 = nil
			} else {
				l2 = l2.Next
			}
		}
	}
	return firstNode
}
