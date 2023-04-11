package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int]()
	assert.True(t, stack.IsEmpty(), "Expected stack to be empty")

	stack.Push(1)
	assert.False(t, stack.IsEmpty(), "Expected stack not to be empty")
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, 2, stack.Peek(), "Expected Peek to return the top item on the stack")
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	res := stack.Pop()

	assert.Equal(t, 2, res, "Expected Pop to return the top item on the stack")
	assert.Equal(t, 1, len(stack.contents), "Expected stack to have one less item after Pop")
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)

	assert.False(t, stack.IsEmpty(), "Expected stack not to be empty after Push")
	assert.Equal(t, 1, len(stack.contents), "Expected stack to have one item after Push")
}
