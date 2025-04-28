# queues [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/queues.svg)](https://pkg.go.dev/github.com/byExist/queues) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/queues)](https://goreportcard.com/report/github.com/byExist/queues)

## What is "queues"?

queues is a lightweight generic queue package written in Go. It uses dynamic resizing and a circular buffer internally to provide efficient and fast queue operations like Enqueue and Dequeue.

## Features

- Supports generic types
- Fast Enqueue/Dequeue operations using a circular buffer
- Automatically grows capacity when needed
- Can be cleared and reused
- Provides a queue copy function
- Supports iteration over values using iter.Seq

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

## Usage

The queues package makes it easy to create and manage lightweight queues. You can quickly perform operations like adding (Enqueue), removing (Dequeue), peeking at the front element (Peek), and iterating over all elements (Values).

Internally, it optimizes performance using a circular buffer and dynamically expands memory when necessary.

## API Overview

### Constructors

- `New[T]() *Queue[T]`
- `NewWithCapacity[T](capacity int) *Queue[T]`
- `Collect[T](seq iter.Seq[T]) *Queue[T]`

### Core Methods

- `Enqueue(q *Queue[T], item T)`
- `Dequeue(q *Queue[T]) (T, bool)`
- `Peek(q *Queue[T]) (T, bool)`
- `Clear(q *Queue[T])`
- `Copy(q *Queue[T]) *Queue[T]`
- `Values(q *Queue[T]) iter.Seq[T]`
- `Len(q *Queue[T]) int`

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
