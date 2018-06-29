package queuewithtwostack

import (
	"github.com/ZhangGuangxu/stack"
)

// QueueWithTwoStack implements a queue with two stack.
type QueueWithTwoStack struct {
	dataSk *stack.Stack
	swapSk *stack.Stack
}

// NewQueueWithTwoStack returns a QueueWithTwoStack instance.
func NewQueueWithTwoStack() *QueueWithTwoStack {
	return &QueueWithTwoStack{
		dataSk: stack.NewStack(),
		swapSk: stack.NewStack(),
	}
}

// PushBack pushes an item at the back of the queue.
func (q *QueueWithTwoStack) PushBack(x interface{}) {
	d := q.dataSk
	if d.IsEmpty() {
		d.Push(x)
		return
	}

	s := q.swapSk
	for !d.IsEmpty() {
		v, _ := d.Pop()
		s.Push(v)
	}
	d.Push(x)
	for !s.IsEmpty() {
		v, _ := s.Pop()
		d.Push(v)
	}
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueWithTwoStack) IsEmpty() bool {
	return q.dataSk.IsEmpty()
}

// PopFront pops an item from the front of the queue.
func (q *QueueWithTwoStack) PopFront() (interface{}, error) {
	return q.dataSk.Pop()
}

// GetFront gets an item from the front of the queue.
func (q *QueueWithTwoStack) GetFront() (interface{}, error) {
	return q.dataSk.Top()
}
