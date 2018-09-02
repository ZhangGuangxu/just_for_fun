package zgxutil

import (
    "errors"
)

// ErrEmptyQueue tells the queue is empty
var ErrEmptyQueue = errors.New("empty queue")

// QueueWithMaxValue 在push和pop后都能追踪当前最大值的队列
type QueueWithMaxValue struct {
    stackA *stackWithMaxValueTrace
    stackB *stackWithMaxValueTrace
}

// NewQueueWithMaxValue creates a new instance of QueueWithMaxValue
func NewQueueWithMaxValue() *QueueWithMaxValue {
    return &QueueWithMaxValue{
        stackA: NewStackWithMaxValueTrace(),
        stackB: NewStackWithMaxValueTrace(),
    }
}

// PushTail pushes one item to the tail of the queue.
func (q *QueueWithMaxValue) PushTail(x int) {
    q.stackB.Push(x)
}

// PopHead pops one item from the head of the queue.
func (q *QueueWithMaxValue) PopHead() (int, error) {
    if q.IsEmpty() {
        return 0, ErrEmptyQueue
    }

    if q.stackA.IsEmpty() {
        for !q.stackB.IsEmpty() {
            b, _ := q.stackB.Pop()
            q.stackA.Push(b)
        }
    }
    return q.stackA.Pop()
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Max tells the max value of the queue.
func (q *QueueWithMaxValue) Max() (int, error) {
    if q.IsEmpty() {
        return 0, ErrEmptyQueue
    }

    ma, _ := q.stackA.Max()
    mb, _ := q.stackB.Max()
    return max(ma, mb), nil
}

// IsEmpty returns true if the queue is empty, otherwise returns false
func (q *QueueWithMaxValue) IsEmpty() bool {
    return q.stackA.IsEmpty() && q.stackB.IsEmpty()
}
