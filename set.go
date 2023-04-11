package collection

type Set[T comparable] struct {
	contents map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		contents: make(map[T]struct{}),
	}
}

func (set *Set[T]) Add(item T) {
	set.contents[item] = struct{}{}
}

func (set *Set[T]) Remove(item T) {
	delete(set.contents, item)
}

func (set *Set[T]) Contains(item T) bool {
	_, ok := set.contents[item]
	return ok
}

func (set Set[T]) Traversal(delegate func(T)) {
	for k := range set.contents {
		delegate(k)
	}
}

func (set Set[T]) Length() int {
	return len(set.contents)
}
