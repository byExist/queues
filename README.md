# queues [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/queues.svg)](https://pkg.go.dev/github.com/byExist/queues) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/queues)](https://goreportcard.com/report/github.com/byExist/queues)

A simple, allocation-efficient FIFO queue implementation in Go.

The `queues` package provides a dynamically resizing, ring-buffer queue supporting enqueue, dequeue, and peek operations. It is designed for high performance in streaming or producer-consumer scenarios with minimal allocations.

---

## ✨ Features

- ✅ Dynamically resizing ring buffer queue
- ✅ Efficient enqueue and dequeue operations
- ✅ Peek at front element without removal
- ✅ Supports generic types with Go 1.18+ generics
- ❌ Not thread-safe (no synchronization)
- ❌ No priority or indexed access

---

## 🧱 Example

```go
package main

import (
	"fmt"
	"github.com/byExist/queues"
)

func main() {
	q := queues.New[int]()

	queues.Enqueue(q, 10)
	queues.Enqueue(q, 20)
	queues.Enqueue(q, 30)

	for v := range queues.Values(q) {
		fmt.Println(v)
	}
}
```

---

## 📚 Use When

- You need FIFO data structure
- You want allocation-efficient buffered queues
- You process items in a streaming or producer-consumer model

---

## 🚫 Avoid If

- You need concurrent access (not thread-safe)
- You want priority-based scheduling or indexed random access

---

## 🔍 API

| Function              | Description                          |
|-----------------------|--------------------------------------|
| `New[T]()`            | Create a new empty queue              |
| `NewWithCapacity[T](capacity int)` | Create a new queue with initial capacity |
| `Collect(seq)`        | Build a queue from an iterator       |
| `Clone()`             | Return a shallow copy of the queue   |
| `Clear()`             | Remove all elements from the queue   |
| `Enqueue(q *Queue[T], item T)` | Add item to the back of the queue    |
| `Dequeue(q *Queue[T]) (T, bool)` | Remove and return front item         |
| `Peek(q *Queue[T]) (T, bool)`    | Return front item without removal    |
| `Len() int`           | Number of elements in queue          |
| `Values()`            | Iterate over queue elements          |

### Methods

| Method                | Description                          |
|------------------------|--------------------------------------|
| `q.String()`          | Returns a string representation       |
| `q.MarshalJSON()` / `q.UnmarshalJSON()` | JSON serialization support |

---

## 🪪 License

MIT License. See [LICENSE](LICENSE).
