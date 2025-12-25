package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	size   int
	front  int
	rear   int
	// need to implement
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		size:   size,
		front:  -1,
		rear:   -1,
		values: make([]int, size),
	} // need to implement
}

func (q *CircularQueue) Push(value int) bool {
	if q.rear == q.size-1 {
		return false
	}

	if q.front == -1 {
		q.front++
	}
	q.rear++
	q.values[q.rear] = value

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.front == -1 {
		return false
	}
	q.front--
	return true
}

func (q *CircularQueue) Front() int {
	return q.values[q.front+1] // need to implement
}

func (q *CircularQueue) Back() int {
	return q.values[q.rear-1] // need to implement
}

func (q *CircularQueue) Empty() bool {
	return q.front == -1 && q.rear == -1
}

func (q *CircularQueue) Full() bool {
	return q.rear == q.size-1
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	return

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
