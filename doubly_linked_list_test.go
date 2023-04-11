package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	/* ---------- Test AddFirst, AddLast, Head, Tail, GetSize, IsEmpty ---------- */
	dll := NewDoublyLinkedList[int]()
	assert.True(t, dll.IsEmpty())
	assert.Nil(t, dll.Head())
	assert.Nil(t, dll.Tail())
	assert.Equal(t, 0, dll.GetSize())

	dll.AddFirst(1)
	assert.False(t, dll.IsEmpty())
	assert.NotNil(t, dll.Head())
	assert.NotNil(t, dll.Tail())
	assert.Equal(t, 1, dll.GetSize())

	dll.AddLast(2)
	assert.Equal(t, 2, dll.GetSize())

	/* ------------------ Test RemoveFirst, RemoveLast, Remove ------------------ */
	value1 := dll.RemoveFirst()
	assert.Equal(t, 1, value1)
	assert.Equal(t, 1, dll.GetSize())

	value2 := dll.RemoveLast()
	assert.Equal(t, 2, value2)
	assert.Equal(t, 0, dll.GetSize())

	dll.AddFirst(3)
	node := dll.Search(3)
	assert.NotNil(t, node)

	value3 := dll.Remove(node)
	assert.Equal(t, 3, value3)
	assert.Equal(t, 0, dll.GetSize())

	/* ------------------------------- Test Clear ------------------------------- */
	dll.AddFirst(4)
	dll.AddLast(5)
	dll.Clear()
	assert.True(t, dll.IsEmpty())
	assert.Nil(t, dll.Head())
	assert.Nil(t, dll.Tail())
	assert.Equal(t, 0, dll.GetSize())
}

func TestDoublyLinkedList_Panic(t *testing.T) {
	/* ------------------------- Test RemoveFirst panic ------------------------- */
	dll := NewDoublyLinkedList[int]()
	assert.Panics(t, func() {
		dll.RemoveFirst()
	})

	/* -------------------------- Test RemoveLast panic ------------------------- */
	assert.Panics(t, func() {
		dll.RemoveLast()
	})

	/* ---------------------------- Test Remove panic --------------------------- */
	dll.Clear()
	assert.Panics(t, func() {
		dll.Remove(nil)
	})
}

func TestDoublyLinkedList_Search(t *testing.T) {
	/* ------------------------------- Test Search ------------------------------ */
	dll := NewDoublyLinkedList[int]()
	dll.AddFirst(1)
	dll.AddLast(2)
	dll.AddLast(3)

	node := dll.Search(2)
	assert.NotNil(t, node)
	assert.Equal(t, 2, node.GetValue())

	node = dll.Search(4)
	assert.Nil(t, node)
}
