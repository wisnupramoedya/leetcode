// problem TLE

package kth_largest_element_in_an_array

type Item struct {
	Value int
	Next  *Item
}

type PriorityQueue struct {
	Biggest *Item
}

func (h *PriorityQueue) Push(pushVal int) {
	thisItem := h.Biggest
	var prevItem *Item = nil

	for {
		if thisItem == nil {
			newItem := &Item{
				Value: pushVal,
				Next:  nil,
			}
			thisItem = newItem

			if h.Biggest == nil {
				h.Biggest = thisItem
			}
			if prevItem != nil && prevItem.Next == nil {
				prevItem.Next = thisItem
			}
			break
		}

		if pushVal >= thisItem.Value {
			newItem := &Item{
				Value: thisItem.Value,
				Next:  thisItem.Next,
			}

			thisItem.Value = pushVal
			thisItem.Next = newItem
			break
		}

		if pushVal < thisItem.Value {
			prevItem = thisItem
			thisItem = thisItem.Next
			continue
		}

		break
	}
}

func (h *PriorityQueue) Find(idx int) int {
	thisItem := h.Biggest
	i := 0
	for {
		if i == idx {
			break
		}

		thisItem = thisItem.Next
		i++
	}

	return thisItem.Value
}

func findKthLargest2(nums []int, k int) int {
	h := &PriorityQueue{Biggest: nil}
	for _, num := range nums {
		h.Push(num)
	}

	return h.Find(k - 1)
}
