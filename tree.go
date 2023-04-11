package collection

// TreeNode represents a node in a generic tree.
type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

// NewTreeNode creates a new TreeNode with the given value.
func NewTreeNode[T any](value T) *TreeNode[T] {
	return &TreeNode[T]{
		Value:    value,
		Children: []*TreeNode[T]{},
	}
}

// AddChild adds a child node to the current node.
func (n *TreeNode[T]) AddChild(child *TreeNode[T]) {
	n.Children = append(n.Children, child)
}

// DepthFirstTraversal performs a depth-first traversal of the tree rooted at the current node and returns a slice of values.
func (n *TreeNode[T]) DepthFirstTraversal() []T {
	var result []T
	result = append(result, n.Value)
	for _, child := range n.Children {
		result = append(result, child.DepthFirstTraversal()...)
	}
	return result
}
