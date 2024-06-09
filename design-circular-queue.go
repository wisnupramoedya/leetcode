package leetcode

type MyCircularQueue struct {
	Value []int
	First int
	Last  int
	Size  int
}

func Constructor(k int) MyCircularQueue {
	var myCircularQueue = MyCircularQueue{
		Value: make([]int, k),
		First: -1,
		Last:  -1,
		Size:  k,
	}
	return myCircularQueue
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.First == -1 {
		this.First = 0
	}
	this.Last = (this.Last + 1) % this.Size
	this.Value[this.Last] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.Last == this.First {
		this.First, this.Last = -1, -1
		return true
	}

	this.First = (this.Size + this.First + 1) % this.Size
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Value[this.First]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Value[this.Last]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.First == -1 && this.Last == -1 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.IsEmpty() {
		return false
	}

	if this.First == (this.Last+1)%this.Size {
		return true
	}

	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
