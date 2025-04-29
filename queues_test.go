package queues_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/byExist/queues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewQueue(t *testing.T) {
	q := queues.New[int]()
	require.NotNil(t, q)
	assert.Equal(t, 0, queues.Len(q))
}

func TestQueueDifferentTypes(t *testing.T) {
	intQ := queues.New[int]()
	queues.Enqueue(intQ, 42)
	v, ok := queues.Dequeue(intQ)
	require.True(t, ok)
	assert.Equal(t, 42, v)

	stringQ := queues.New[string]()
	queues.Enqueue(stringQ, "hello")
	s, ok := queues.Dequeue(stringQ)
	require.True(t, ok)
	assert.Equal(t, "hello", s)

	type Person struct {
		Name string
		Age  int
	}

	personQ := queues.New[Person]()
	queues.Enqueue(personQ, Person{Name: "Alice", Age: 30})
	p, ok := queues.Dequeue(personQ)
	require.True(t, ok)
	assert.Equal(t, Person{Name: "Alice", Age: 30}, p)
}

func TestCollect(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	seq := func(yield func(int) bool) {
		for _, v := range data {
			if !yield(v) {
				return
			}
		}
	}
	q := queues.Collect(seq)
	require.Equal(t, len(data), queues.Len(q))
	values := slices.Collect(queues.Values(q))
	assert.Equal(t, data, values)
}

func TestEnqueueSingle(t *testing.T) {
	q := queues.New[int]()
	queues.Enqueue(q, 10)
	require.Equal(t, 1, queues.Len(q))

	v, ok := queues.Peek(q)
	require.True(t, ok)
	assert.Equal(t, 10, v)
}

func TestEnqueueMultiple(t *testing.T) {
	q := queues.New[int]()
	for i := range 5 {
		queues.Enqueue(q, i)
	}
	require.Equal(t, 5, queues.Len(q))

	values := slices.Collect(queues.Values(q))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, values)
}

func TestEnqueueGrow(t *testing.T) {
	q := queues.New[int]()
	queues.Enqueue(q, 1)
	queues.Enqueue(q, 2)
	queues.Enqueue(q, 3)
	assert.Equal(t, 3, queues.Len(q))
}

func TestPeekEmpty(t *testing.T) {
	q := queues.New[int]()
	_, ok := queues.Peek(q)
	assert.False(t, ok)
}

func TestPeekNonEmpty(t *testing.T) {
	q := queues.New[int]()
	queues.Enqueue(q, 99)
	v, ok := queues.Peek(q)
	require.True(t, ok)
	assert.Equal(t, 99, v)

	// Peek does not remove
	require.Equal(t, 1, queues.Len(q))
}

func TestDequeueEmpty(t *testing.T) {
	q := queues.New[int]()
	_, ok := queues.Dequeue(q)
	assert.False(t, ok)
}

func TestDequeueSingle(t *testing.T) {
	q := queues.New[int]()
	queues.Enqueue(q, 42)
	v, ok := queues.Dequeue(q)
	require.True(t, ok)
	assert.Equal(t, 42, v)
	assert.Equal(t, 0, queues.Len(q))
}

func TestDequeueMultiple(t *testing.T) {
	q := queues.New[int]()
	for i := 1; i <= 5; i++ {
		queues.Enqueue(q, i)
	}
	for i := 1; i <= 5; i++ {
		v, ok := queues.Dequeue(q)
		require.True(t, ok)
		assert.Equal(t, i, v)
	}
	assert.Equal(t, 0, queues.Len(q))
}

func TestLenBehavior(t *testing.T) {
	q := queues.New[int]()
	assert.Equal(t, 0, queues.Len(q))

	queues.Enqueue(q, 10)
	assert.Equal(t, 1, queues.Len(q))

	queues.Enqueue(q, 20)
	assert.Equal(t, 2, queues.Len(q))

	_, _ = queues.Dequeue(q)
	assert.Equal(t, 1, queues.Len(q))
}

func TestValuesEmpty(t *testing.T) {
	q := queues.New[int]()
	values := slices.Collect(queues.Values(q))
	assert.Empty(t, values)
}

func TestValuesMultiple(t *testing.T) {
	q := queues.New[int]()
	expected := []int{10, 20, 30}
	for _, v := range expected {
		queues.Enqueue(q, v)
	}
	result := slices.Collect(queues.Values(q))
	assert.Equal(t, expected, result)
}

func TestClearBehavior(t *testing.T) {
	q := queues.New[int]()
	for i := range 5 {
		queues.Enqueue(q, i)
	}
	require.Equal(t, 5, queues.Len(q))

	queues.Clear(q)
	assert.Equal(t, 0, queues.Len(q))

	// Enqueue/Dequeue after Clear
	queues.Enqueue(q, 100)
	v, ok := queues.Dequeue(q)
	require.True(t, ok)
	assert.Equal(t, 100, v)
}

func TestCopyBehavior(t *testing.T) {
	q := queues.New[int]()
	for i := 1; i <= 3; i++ {
		queues.Enqueue(q, i)
	}

	copied := queues.Copy(q)
	require.Equal(t, queues.Len(q), queues.Len(copied))

	origValues := slices.Collect(queues.Values(q))
	copyValues := slices.Collect(queues.Values(copied))
	assert.Equal(t, origValues, copyValues)

	// Modify original and ensure copy is unaffected
	queues.Enqueue(q, 99)
	assert.NotEqual(t, queues.Len(q), queues.Len(copied))
}

func TestCircularBufferWrapAround(t *testing.T) {
	q := queues.New[int]()

	queues.Enqueue(q, 1)
	queues.Enqueue(q, 2)
	queues.Enqueue(q, 3)

	v, ok := queues.Dequeue(q)
	require.True(t, ok)
	assert.Equal(t, 1, v)

	queues.Enqueue(q, 4)

	values := slices.Collect(queues.Values(q))
	assert.Equal(t, []int{2, 3, 4}, values)
}

func TestLargeDataHandling(t *testing.T) {
	q := queues.New[int]()
	n := 10000
	for i := range n {
		queues.Enqueue(q, i)
	}
	assert.Equal(t, n, queues.Len(q))

	for i := range n {
		v, ok := queues.Dequeue(q)
		require.True(t, ok)
		assert.Equal(t, i, v)
	}
	assert.Equal(t, 0, queues.Len(q))
}

func ExampleNew() {
	q := queues.New[int]()
	queues.Enqueue(q, 10)
	v, _ := queues.Dequeue(q)
	fmt.Println(v)
	// Output: 10
}

func ExampleCollect() {
	seq := slices.Values([]int{1, 2, 3})
	q := queues.Collect(seq)
	for v := range queues.Values(q) {
		fmt.Print(v)
	}
	// Output: 123
}

func ExampleCopy() {
	q := queues.New[int]()
	queues.Enqueue(q, 1)
	queues.Enqueue(q, 2)

	copied := queues.Copy(q)
	for v := range queues.Values(copied) {
		fmt.Print(v)
	}
	// Output: 12
}

func ExamplePeek() {
	q := queues.New[int]()
	queues.Enqueue(q, 99)
	v, _ := queues.Peek(q)
	fmt.Println(v)
	// Output: 99
}

func ExampleEnqueue() {
	q := queues.New[int]()
	queues.Enqueue(q, 5)
	v, _ := queues.Peek(q)
	fmt.Println(v)
	// Output: 5
}

func ExampleDequeue() {
	q := queues.New[int]()
	queues.Enqueue(q, 7)
	v, _ := queues.Dequeue(q)
	fmt.Println(v)
	// Output: 7
}

func ExampleLen() {
	q := queues.New[int]()
	fmt.Println(queues.Len(q))
	queues.Enqueue(q, 1)
	fmt.Println(queues.Len(q))
	// Output:
	// 0
	// 1
}

func ExampleValues() {
	q := queues.New[int]()
	queues.Enqueue(q, 1)
	queues.Enqueue(q, 2)
	for v := range queues.Values(q) {
		fmt.Print(v)
	}
	// Output: 12
}

func ExampleClear() {
	q := queues.New[int]()
	queues.Enqueue(q, 1)
	queues.Clear(q)
	fmt.Println(queues.Len(q))
	// Output: 0
}
