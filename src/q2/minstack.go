package minstack

import (
	"errors"
	"github.com/ZhangGuangxu/stack"
)

// 实现一个带有min功能的栈，要求min、push、pop的时间复杂度都是O(1)

// ErrType tells us some type is wrong.
var ErrType = errors.New("type error")

// MinStack has Min func.
type MinStack struct {
	dataSk *stack.Stack
	minSk *stack.Stack
}

// NewMinStack creates a new MinStack instance.
func NewMinStack() *MinStack {
	return &MinStack{
		dataSk: stack.NewStack(),
		minSk: stack.NewStack(),
	}
}

// Push pushes an item onto the top of the stack.
func (s *MinStack) Push(x int) error {
	d := s.dataSk
	empty := d.IsEmpty()
	d.Push(x)
	m := s.minSk

	if empty {
		m.Push(x)
	} else {
		v, _ := m.Top()
		tv := 0
		ok := false
		if tv, ok = v.(int); !ok {
			return ErrType
		}
		if x <= tv {
			m.Push(x)
		}
	}

	return nil
}

// Pop pops an item from the top of the stack.
func (s *MinStack) Pop() (int, error) {
	d := s.dataSk
	dv, err := d.Pop()
	if err != nil {
		return 0, err
	}

	m := s.minSk
	tv, err := m.Top()
	if err != nil {
		return 0, err
	}

	ok := false
	dv2 := 0
	if dv2, ok = dv.(int); !ok {
		return 0, ErrType
	}
	tv2 := 0
	if tv2, ok = tv.(int); !ok {
		return 0, ErrType
	}
	if dv2 == tv2 {
		if _, err := m.Pop(); err != nil {
			return 0, err
		}
	}

	return dv2, nil
}

// Min returns min value of the stack.
func (s *MinStack) Min() (int, error) {
	v, err := s.minSk.Top()
	if err != nil {
		return 0, err
	}

	tv := 0
	ok := false
	if tv, ok = v.(int); !ok {
		return 0, ErrType
	}

	return tv, nil
}
