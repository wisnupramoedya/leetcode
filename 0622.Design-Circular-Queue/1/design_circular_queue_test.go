package design_circular_queue

import (
	"github.com/wisnupramoedya/leetcode/utils"
	"testing"
)

func run(commands []string, args [][]int) []interface{} {
	var results []interface{}
	var myCircularQueue MyCircularQueue

	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "MyCircularQueue":
			myCircularQueue = Constructor(args[i][0])
			results = append(results, nil)
		case "enQueue":
			results = append(results, myCircularQueue.EnQueue(args[i][0]))
		case "deQueue":
			results = append(results, myCircularQueue.DeQueue())
		case "Front":
			results = append(results, myCircularQueue.Front())
		case "Rear":
			results = append(results, myCircularQueue.Rear())
		case "isEmpty":
			results = append(results, myCircularQueue.IsEmpty())
		case "isFull":
			results = append(results, myCircularQueue.IsFull())
		}
	}

	return results
}

func Test622_1_1(t *testing.T) {
	commands := []string{"MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"}
	args := [][]int{{3}, {1}, {2}, {3}, {4}, {}, {}, {}, {4}, {}}

	output := []interface{}{nil, true, true, true, false, 3, true, true, true, 4}

	result := run(commands, args)

	if !utils.IsArraySame(output, result) {
		t.Errorf("It is not the same array")
	}

}
