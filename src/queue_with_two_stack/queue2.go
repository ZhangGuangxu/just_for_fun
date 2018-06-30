package queuewithtwostack

import (
	"github.com/ZhangGuangxu/stack"
)

type queue2 struct {
	skPush *stack.Stack
	skPop  *stack.Stack
}

func newQueue2() *queue2 {
	return &queue2{
		skPush: stack.NewStack(),
		skPop:  stack.NewStack(),
	}
}

func (q *queue2) pushBack(x interface{}) {
	skPush := q.skPush
	skPop := q.skPop
	for !skPop.IsEmpty() {
		v, _ := skPop.Pop()
		skPush.Push(v)
	}
	skPush.Push(x)
}

func (q *queue2) popFront() (interface{}, error) {
	skPop := q.skPop
	if !skPop.IsEmpty() {
		return skPop.Pop()
	}

	skPush := q.skPush
	for !skPush.IsEmpty() {
		v, _ := skPush.Pop()
		skPop.Push(v)
	}
	return skPop.Pop()
}

func (q *queue2) getFront() (interface{}, error) {
	skPop := q.skPop
	if !skPop.IsEmpty() {
		return skPop.Top()
	}

	skPush := q.skPush
	for !skPush.IsEmpty() {
		v, _ := skPush.Pop()
		skPop.Push(v)
	}
	return skPop.Top()
}
