# queues [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/queues.svg)](https://pkg.go.dev/github.com/byExist/queues) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/queues)](https://goreportcard.com/report/github.com/byExist/queues)

## What is "queues"?

`queues` is a lightweight generic queue package written in Go. It provides efficient queue operations like Enqueue, Dequeue, and Peek, and supports iteration over elements using `iter.Seq`. Internally, it uses a circular buffer with dynamic resizing to ensure fast performance even as elements are added and removed.

This package supports generic types, automatically grows capacity when needed, and allows for queue cloning, clearing, and reuse. It also provides convenient string and JSON representations of queues for easier debugging and integration.


## Installation

To install, use the following command:

```bash
go get github.com/byExist/queues
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/byExist/queues"
)

func main() {
	q := queues.New[int]()

	// Enqueue elements
	queues.Enqueue(q, 1)
	queues.Enqueue(q, 2)
	queues.Enqueue(q, 3)

	// Peek at the front element
	v, ok := queues.Peek(q)
	if ok {
		fmt.Println("Peek:", v)
	}

	// Dequeue elements
	for {
		v, ok := queues.Dequeue(q)
		if !ok {
			break
		}
		fmt.Println("Dequeue:", v)
	}

	// Check length
	fmt.Println("Length:", queues.Len(q))
}
```

```output
Peek: 1
Dequeue: 1
Dequeue: 2
Dequeue: 3
Length: 0
```


## API Overview

### Constructors

| Function                            | Description                         | Time Complexity |
|-------------------------------------|-------------------------------------|-----------------|
| `New[T]()`                          | Create a new empty queue            | O(1)            |
| `Collect[T](seq iter.Seq[T])`       | Build a queue from an iterator      | O(n)            |

### Operations

| Function                            | Description                         | Time Complexity |
|-------------------------------------|-------------------------------------|-----------------|
| `Enqueue(q *Queue[T], item T)`      | Add an element to the end           | Amortized O(1)  |
| `Dequeue(q *Queue[T]) (T, bool)`    | Remove and return the front element | O(1)            |
| `Peek(q *Queue[T]) (T, bool)`       | Peek at the front element           | O(1)            |
| `Clear(q *Queue[T])`                | Remove all elements                 | O(1)            |
| `Clone(q *Queue[T]) *Queue[T]`      | Create a shallow copy of the queue  | O(n)            |

### Introspection

| Function                            | Description                         | Time Complexity |
|-------------------------------------|-------------------------------------|-----------------|
| `Len(q *Queue[T]) int`              | Return the number of elements       | O(1)            |

### Iteration

| Function                            | Description                         | Time Complexity |
|-------------------------------------|-------------------------------------|-----------------|
| `Values(q *Queue[T]) iter.Seq[T]`   | Get an iterator over the queue      | O(1)            |

### Methods

| Method                              | Description                         | Time Complexity |
|-------------------------------------|-------------------------------------|-----------------|
| `(*Queue[T]) String() string`       | Return a string representation      | O(n)            |
| `(*Queue[T]) MarshalJSON() ([]byte, error)` | Serialize the queue to JSON   | O(n)            |
| `(*Queue[T]) UnmarshalJSON([]byte) error`   | Parse a JSON array into a queue | O(n)            |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
