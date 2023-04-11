package collection

type Queue[T any] struct {
	contents []T
}

func New[T any]() Queue[T] {
	return Queue[T]{
		contents: []T{},
	}
}

func NewQueueFromSlice[T any](items []T) Queue[T] {
	if items == nil {
		return Queue[T]{
			contents: []T{},
		}
	}
	return Queue[T]{
		contents: items,
	}
}

func (q *Queue[T]) Enqueue(element T) {
	q.contents = append(q.contents, element)
}

// please make sure the queue is not empty or the method will raise a panic.
func (q *Queue[T]) Dequeue() T {
	res := q.Peek()
	q.contents = q.contents[1:]
	return res
}

// please make sure the queue is not empty or the method will raise a panic.
func (q *Queue[T]) Peek() T {
	return q.contents[0]
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.contents) == 0
}
