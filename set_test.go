package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	set := NewSet[int]()
	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Length())
}

func TestAdd(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.Equal(t, 3, set.Length())
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
}

func TestRemove(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Remove(2)
	assert.Equal(t, 2, set.Length())
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
	assert.True(t, set.Contains(3))

	set.Remove(1)
	assert.Equal(t, 1, set.Length())
	assert.False(t, set.Contains(1))
	assert.False(t, set.Contains(2))
	assert.True(t, set.Contains(3))
}

func TestContains(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
	assert.False(t, set.Contains(4))
}

func TestTraversal(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	sum := 0
	set.Traversal(func(item int) {
		sum += item
	})
	assert.Equal(t, 6, sum)
}

func TestLength(t *testing.T) {
	set := NewSet[int]()
	assert.Equal(t, 0, set.Length())

	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.Equal(t, 3, set.Length())

	set.Remove(2)
	assert.Equal(t, 2, set.Length())
}
