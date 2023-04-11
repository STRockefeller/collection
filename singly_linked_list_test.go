package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSinglyLinkedListNode(t *testing.T) {
	node := NewSinglyLinkedListNode(42)
	assert.NotNil(t, node)
	assert.Equal(t, 42, node.GetValue())
	assert.Nil(t, node.GetNext())
}

func TestNewSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[struct{}]()
	assert.NotNil(t, list)
	assert.True(t, list.IsEmpty())
	assert.Nil(t, list.Head())
}

func TestAddNode(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	node1 := NewSinglyLinkedListNode(1)
	list.AddNode(node1)

	assert.False(t, list.IsEmpty())
	assert.Equal(t, node1, list.Head())

	node2 := NewSinglyLinkedListNode(2)
	list.AddNode(node2)

	assert.Equal(t, node2, node1.GetNext())
}

func TestDeleteNode(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	node1 := NewSinglyLinkedListNode(1)
	node2 := NewSinglyLinkedListNode(2)
	node3 := NewSinglyLinkedListNode(3)

	list.AddNode(node1)
	list.AddNode(node2)
	list.AddNode(node3)

	assert.Equal(t, node1, list.Head())

	// Delete node2
	result := list.DeleteNode(node2)
	assert.True(t, result)
	assert.Equal(t, node3, node1.GetNext())

	// Delete node1
	result = list.DeleteNode(node1)
	assert.True(t, result)
	assert.Equal(t, node3, list.Head())

	// Delete node3
	result = list.DeleteNode(node3)
	assert.True(t, result)
	assert.True(t, list.IsEmpty())

	// Delete non-existing node
	node4 := NewSinglyLinkedListNode(4)
	result = list.DeleteNode(node4)
	assert.False(t, result)
}

func TestSearch(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	node1 := NewSinglyLinkedListNode(1)
	node2 := NewSinglyLinkedListNode(2)
	node3 := NewSinglyLinkedListNode(3)

	list.AddNode(node1)
	list.AddNode(node2)
	list.AddNode(node3)

	result := list.Search(2)
	assert.Equal(t, node2, result)

	result = list.Search(4)
	assert.Nil(t, result)
}

func TestSinglyLinkedListNodeMethods(t *testing.T) {
	node := NewSinglyLinkedListNode("test")
	assert.Equal(t, "test", node.GetValue())

	node2 := NewSinglyLinkedListNode("42")
	node.SetNext(node2)
	assert.Equal(t, node2, node.GetNext())
}
