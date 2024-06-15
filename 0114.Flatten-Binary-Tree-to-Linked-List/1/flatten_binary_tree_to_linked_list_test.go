package flatten_binary_tree_to_linked_list

import (
	"github.com/wisnupramoedya/leetcode/structures"
	"testing"
)

func run(root, output []interface{}) (*TreeNode, *TreeNode) {
	rootTreeNode := structures.SetArrayIntNil2TreeNode(root)
	flatten(rootTreeNode)
	outputTreeNode := structures.SetArrayIntNil2TreeNode(output)
	return rootTreeNode, outputTreeNode
}

func Test114_1_1(t *testing.T) {
	root := []interface{}{1, 2, 5, 3, 4, nil, 6}
	output := []interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6}

	rootTreeNode, outputTreeNode := run(root, output)

	if !structures.IsTreeNodeEqual(rootTreeNode, outputTreeNode, 0) {
		t.Errorf("It is not the same tree node")
	}
}

func Test114_1_2(t *testing.T) {
	var root []interface{}
	var output []interface{}

	rootTreeNode, outputTreeNode := run(root, output)

	if !structures.IsTreeNodeEqual(rootTreeNode, outputTreeNode, 0) {
		t.Errorf("It is not the same tree node")
	}
}

func Test114_1_3(t *testing.T) {
	root := []interface{}{0}
	output := []interface{}{0}

	rootTreeNode, outputTreeNode := run(root, output)

	if !structures.IsTreeNodeEqual(rootTreeNode, outputTreeNode, 0) {
		t.Errorf("It is not the same tree node")
	}
}
