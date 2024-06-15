package flatten_binary_tree_to_linked_list

/*
Runtime: 0 ms
Memory Usage: 3.1 MB
*/

import (
	"github.com/wisnupramoedya/leetcode/structures"
)

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	newStack := *s
	length := len(newStack)
	value := newStack[length-1]
	*s = newStack[0 : length-1]
	//println("pop", value)
	return value
}

func traceBinaryTree(stack *Stack, node *TreeNode) {
	if node == nil {
		return
	}

	stack.Push(node)
	//println("push", node.Val)

	if node.Left != nil {
		//println("left", node.Left.Val)
		traceBinaryTree(stack, node.Left)
	}

	if node.Right != nil {
		//println("right", node.Right.Val)
		traceBinaryTree(stack, node.Right)
	}
}

func flatten(root *TreeNode) {
	stack := Stack{}
	traceBinaryTree(&stack, root)
	//println("root", root)

	var rightNode *TreeNode
	for {
		if len(stack) == 0 {
			//root.Left = nil
			//root.Right = rightNode
			break
		}

		node := stack.Pop()
		node.Left = nil
		node.Right = rightNode
		//if rightNode != nil {
		//	println(node, node.Val, node.Left, rightNode, rightNode.Val)
		//} else {
		//	println(node, node.Val, node.Left, rightNode)
		//}

		rightNode = node
	}
}
