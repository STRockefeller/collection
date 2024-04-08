package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.NotNil(t, q)
	assert.Equal(t, 0, len(q.contents))
}

func TestNewQueueFromSlice(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	q := NewQueueFromSlice(items)
	assert.NotNil(t, q)
	assert.Equal(t, len(items), len(q.contents))
	for i, item := range items {
		assert.Equal(t, item, q.contents[i])
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, 3, len(q.contents))
	assert.Equal(t, 1, q.contents[0])
	assert.Equal(t, 2, q.contents[1])
	assert.Equal(t, 3, q.contents[2])
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item := q.Dequeue()
	assert.Equal(t, 1, item)
	assert.Equal(t, 2, len(q.contents))
	assert.Equal(t, 2, q.contents[0])
	assert.Equal(t, 3, q.contents[1])
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.Panics(t, func() {
		q.Dequeue()
	})
}

func TestPeek(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item := q.Peek()
	assert.Equal(t, 1, item)
	assert.Equal(t, 3, len(q.contents))
	assert.Equal(t, 1, q.contents[0])
	assert.Equal(t, 2, q.contents[1])
	assert.Equal(t, 3, q.contents[2])
}

func TestPeekEmptyQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.Panics(t, func() {
		q.Peek()
	})
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue[int]()
	assert.True(t, q.IsEmpty())

	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}
