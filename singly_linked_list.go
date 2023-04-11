package collection

import "reflect"

type SinglyLinkedListNode[T any] struct {
	value T
	next  *SinglyLinkedListNode[T]
}

func (n *SinglyLinkedListNode[T]) GetValue() T {
	return n.value
}

func (n *SinglyLinkedListNode[T]) GetNext() *SinglyLinkedListNode[T] {
	return n.next
}

func (n *SinglyLinkedListNode[T]) SetNext(node *SinglyLinkedListNode[T]) {
	n.next = node
}

func NewSinglyLinkedListNode[T any](value T) *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		value: value,
		next:  nil,
	}
}

type SinglyLinkedList[T any] struct {
	head *SinglyLinkedListNode[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *SinglyLinkedList[T]) Head() *SinglyLinkedListNode[T] {
	return l.head
}

func (l *SinglyLinkedList[T]) AddNode(node *SinglyLinkedListNode[T]) {
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.GetNext() != nil {
			current = current.GetNext()
		}
		current.SetNext(node)
	}
}

func (l *SinglyLinkedList[T]) DeleteNode(node *SinglyLinkedListNode[T]) bool {
	if l.IsEmpty() {
		return false
	}
	if l.head == node {
		l.head = node.GetNext()
		return true
	}
	current := l.head
	for current.GetNext() != nil {
		if current.GetNext() == node {
			current.SetNext(node.GetNext())
			return true
		}
		current = current.GetNext()
	}

	return false
}

func (l *SinglyLinkedList[T]) Search(value T) *SinglyLinkedListNode[T] {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(current.GetValue(), value) {
			return current
		}
		current = current.GetNext()
	}
	return nil
}
