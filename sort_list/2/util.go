package sort_list2

type ListNode struct {
	Val  int
	Next *ListNode
}

func setLinkedList(values []int) *ListNode {
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

func setLinkedListToArrayInt(head *ListNode) []int {
	var values []int
	for {
		if head == nil {
			break
		}

		values = append(values, head.Val)
		head = head.Next
	}
	return values
}
