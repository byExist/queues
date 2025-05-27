package queues

import (
	"encoding/json"
	"fmt"
	"iter"
	"slices"
	"strings"
)

// Queue is a generic, dynamically resizing circular queue.
type Queue[T any] struct {
	items      []T
	head, tail int
	size       int
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	var b strings.Builder
	b.WriteString("Queue{")
	for i := range q.size {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprint(q.items[(q.head+i)%len(q.items)]))
	}
	b.WriteString("}")
	return b.String()
}

// MarshalJSON implements json.Marshaler for Queue.
func (q *Queue[T]) MarshalJSON() ([]byte, error) {
	values := make([]T, q.size)
	for i := range q.size {
		values[i] = q.items[(q.head+i)%len(q.items)]
	}
	return json.Marshal(values)
}

// UnmarshalJSON implements json.Unmarshaler for Queue.
func (q *Queue[T]) UnmarshalJSON(data []byte) error {
	var values []T
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	q.items = make([]T, len(values))
	copy(q.items, values)
	q.head = 0
	if len(q.items) > 0 {
		q.tail = len(values) % len(q.items)
	} else {
		q.tail = 0
	}
	q.size = len(values)
	return nil
}

// New creates a new empty queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// NewWithCapacity creates a new queue with a specified initial capacity.
func NewWithCapacity[T any](capacity int) *Queue[T] {
	if capacity < 0 {
		panic("capacity cannot be negative")
	}
	return &Queue[T]{items: make([]T, 0, capacity)}
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
