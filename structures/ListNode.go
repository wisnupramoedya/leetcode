package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func SetArrayInt2ListNode(values []int) *ListNode {
	var firstListNode, prevListNode *ListNode
	for _, value := range values {
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

func SetListNode2ArrayInt(head *ListNode) []int {
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
