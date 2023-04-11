package collection

import "reflect"

type DoublyLinkedList[T any] struct {
	head *DoublyLinkedListNode[T]
	tail *DoublyLinkedListNode[T]
	size int
}

type DoublyLinkedListNode[T any] struct {
	value T
	prev  *DoublyLinkedListNode[T]
	next  *DoublyLinkedListNode[T]
}

func NewDoublyLinkedListNode[T any](value T) *DoublyLinkedListNode[T] {
	return &DoublyLinkedListNode[T]{
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func (n *DoublyLinkedListNode[T]) GetValue() T {
	return n.value
}

func (n *DoublyLinkedListNode[T]) GetPrev() *DoublyLinkedListNode[T] {
	return n.prev
}

func (n *DoublyLinkedListNode[T]) SetPrev(node *DoublyLinkedListNode[T]) {
	n.prev = node
}

func (n *DoublyLinkedListNode[T]) GetNext() *DoublyLinkedListNode[T] {
	return n.next
}

func (n *DoublyLinkedListNode[T]) SetNext(node *DoublyLinkedListNode[T]) {
	n.next = node
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (dll *DoublyLinkedList[T]) GetSize() int {
	return dll.size
}

func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoublyLinkedList[T]) AddFirst(value T) {
	newNode := NewDoublyLinkedListNode(value)

	if dll.IsEmpty() {
		dll.tail = newNode
	} else {
		dll.head.SetPrev(newNode)
		newNode.SetNext(dll.head)
	}

	dll.head = newNode
	dll.size++
}

func (dll *DoublyLinkedList[T]) AddLast(value T) {
	newNode := NewDoublyLinkedListNode(value)

	if dll.IsEmpty() {
		dll.head = newNode
	} else {
		dll.tail.SetNext(newNode)
		newNode.SetPrev(dll.tail)
	}

	dll.tail = newNode
	dll.size++
}

// will panic if list is empty
func (dll *DoublyLinkedList[T]) RemoveFirst() T {
	if dll.IsEmpty() {
		panic("List is empty")
	}

	value := dll.head.GetValue()

	if dll.GetSize() == 1 {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.GetNext()
		dll.head.SetPrev(nil)
	}

	dll.size--
	return value
}

// will panic if list is empty
func (dll *DoublyLinkedList[T]) RemoveLast() T {
	if dll.IsEmpty() {
		panic("List is empty")
	}

	value := dll.tail.GetValue()

	if dll.GetSize() == 1 {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.GetPrev()
		dll.tail.SetNext(nil)
	}

	dll.size--
	return value
}

// will panic if given node is nil
func (dll *DoublyLinkedList[T]) Remove(node *DoublyLinkedListNode[T]) T {
	if node == nil {
		panic("Node is nil")
	}

	if node == dll.head {
		return dll.RemoveFirst()
	}

	if node == dll.tail {
		return dll.RemoveLast()
	}

	prevNode := node.GetPrev()
	nextNode := node.GetNext()

	prevNode.SetNext(nextNode)
	nextNode.SetPrev(prevNode)

	dll.size--
	return node.GetValue()
}

func (dll *DoublyLinkedList[T]) Clear() {
	for dll.head != nil {
		dll.RemoveFirst()
	}
}

func (dll *DoublyLinkedList[T]) Search(value T) *DoublyLinkedListNode[T] {
	current := dll.head
	for current != nil {
		if reflect.DeepEqual(current.GetValue(), value) {
			return current
		}
		current = current.GetNext()
	}
	return nil
}

func (dll *DoublyLinkedList[T]) Head() *DoublyLinkedListNode[T] {
	return dll.head
}

func (dll *DoublyLinkedList[T]) Tail() *DoublyLinkedListNode[T] {
	return dll.tail
}
