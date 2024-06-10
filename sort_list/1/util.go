package sort_list1

type ListNode struct {
	Val  int
	Next *ListNode
}

func SetLinkedListToArrayInt(head *ListNode) []int {
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
