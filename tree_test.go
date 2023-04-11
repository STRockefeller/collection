package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeNode_AddChild(t *testing.T) {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)

	node1.AddChild(node2)
	node1.AddChild(node3)

	assert.Equal(t, []*TreeNode[int]{node2, node3}, node1.Children, "Expected child nodes to be added to the parent node")
}

func TestTreeNode_DepthFirstTraversal(t *testing.T) {
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)

	node1.AddChild(node2)
	node1.AddChild(node3)
	node2.AddChild(node4)

	result := node1.DepthFirstTraversal()

	assert.Equal(t, []int{1, 2, 4, 3}, result, "Expected depth-first traversal result to be in the correct order")
}

func TestTreeNode_DepthFirstTraversal_Empty(t *testing.T) {
	node := NewTreeNode(1)

	result := node.DepthFirstTraversal()

	assert.Equal(t, []int{1}, result, "Expected depth-first traversal result to contain only the root node value when the tree is empty")
}
