package queues

import (
	"iter"
	"slices"
)

// Queue is a generic, dynamically resizing circular queue.
type Queue[T any] struct {
	items      []T
	head, tail int
	size       int
}

// grow increases the capacity of the queue and reorders elements to maintain sequence.
func (q *Queue[T]) grow() {
	var newCap int
	if len(q.items) < 1024 {
		newCap = 2 * max(1, len(q.items))
	} else {
		newCap = len(q.items) + len(q.items)/4
	}

	newItems := make([]T, newCap)
	for i := range q.size {
		newItems[i] = q.items[(q.head+i)%len(q.items)]
	}

	q.items = newItems
	q.head = 0
	q.tail = q.size
}

// New creates a new empty queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Collect builds a queue from a given sequence of elements.
func Collect[T any](i iter.Seq[T]) *Queue[T] {
	q := New[T]()
	for e := range i {
		Enqueue(q, e)
	}
	return q
}

// Clone creates a new queue with the same elements as the given queue.
func Clone[T any](q *Queue[T]) *Queue[T] {
	return &Queue[T]{
		items: slices.Clone(q.items),
		head:  q.head,
		tail:  q.tail,
		size:  q.size,
	}
}

// Peek returns the front element without removing it. Returns false if the queue is empty.
func Peek[T any](q *Queue[T]) (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}
	return q.items[q.head], true
}

// Enqueue adds an element to the end of the queue.
func Enqueue[T any](q *Queue[T], item T) {
	if q.size == len(q.items) {
		q.grow()
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % len(q.items)
	q.size++
}

// Dequeue removes and returns the front element of the queue. Returns false if the queue is empty.
func Dequeue[T any](q *Queue[T]) (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}
	item := q.items[q.head]
	q.head = (q.head + 1) % len(q.items)
	q.size--
	return item, true
}

// Len returns the number of elements in the queue.
func Len[T any](q *Queue[T]) int {
	return q.size
}

// Values returns a sequence that yields all elements in the queue in order.
func Values[T any](q *Queue[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range q.size {
			if !yield(q.items[(q.head+i)%len(q.items)]) {
				break
			}
		}
	}
}

// Clear removes all elements but keeps the allocated capacity.
func Clear[T any](q *Queue[T]) {
	q.items = make([]T, 0, len(q.items))
	q.head = 0
	q.tail = 0
	q.size = 0
}
