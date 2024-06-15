package structures

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SetArrayIntNil2TreeNode(values []interface{}) *TreeNode {
	maxLen := len(values)
	if maxLen == 0 {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	queue := make([]*TreeNode, 1, maxLen<<1)
	queue[0] = root

	i := 1
	for i < maxLen {
		node := queue[0]
		queue = queue[1:]
		if values[i] != nil && i < maxLen {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if values[i] != nil && i < maxLen {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func IsTreeNodeEqual(firstTreeNode, secondTreeNode *TreeNode, height int) bool {
	if firstTreeNode == nil && secondTreeNode == nil {
		fmt.Printf("In height %d, all node is nil\n", height)
		return true
	}

	if firstTreeNode == nil || secondTreeNode == nil || firstTreeNode.Val != secondTreeNode.Val {
		if firstTreeNode == nil {
			fmt.Printf("In height %d, first node is nil, while second node is %d\n", height, secondTreeNode.Val)
		} else if secondTreeNode == nil {
			fmt.Printf("In height %d, first node is %d, while second node is nil\n", height, firstTreeNode.Val)
		} else {
			fmt.Printf("In height %d, first node is %d, while second node is %d\n", height, firstTreeNode.Val, secondTreeNode.Val)
		}
		return false
	}

	fmt.Printf("Height: %d => %d == %d\n", height, firstTreeNode.Val, secondTreeNode.Val)
	return IsTreeNodeEqual(firstTreeNode.Left, secondTreeNode.Left, height+1) && IsTreeNodeEqual(firstTreeNode.Right, secondTreeNode.Right, height+1)
}
