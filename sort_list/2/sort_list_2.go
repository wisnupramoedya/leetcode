package sort_list2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge(left, right *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = left
			left = left.Next
		} else {
			node.Next = right
			right = right.Next
		}
		node = node.Next
	}
	if left != nil {
		node.Next = left
	}
	if right != nil {
		node.Next = right
	}

	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prevMid, mid, last := head, head, head

	// step 1: set the real mid and last
	for last != nil && last.Next != nil {
		prevMid, mid, last = mid, mid.Next, last.Next.Next
	}
	prevMid.Next = nil

	// step 2: sort the list of head -> prevMid and mid -> last
	listA, listB := sortList(head), sortList(mid)

	return merge(listA, listB)
}
