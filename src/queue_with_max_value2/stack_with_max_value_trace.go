package zgxutil

import (
    "errors"
)

// ErrEmptyStack 表明空栈
var ErrEmptyStack = errors.New("empty stack")
// ErrNoMax 表明没有最大值
var ErrNoMax = errors.New("no max")

// 能追踪每一步push/pop后当前的最大值
type stackWithMaxValueTrace struct {
    items []int
    top int

    link2NextMaxItem []int
    maxItemIndex int
}

// NewStackWithMaxValueTrace creates a new stackWithMaxValueTrace instance
func NewStackWithMaxValueTrace() *stackWithMaxValueTrace {
    return &stackWithMaxValueTrace{
        top: -1,
        maxItemIndex: -1,
    }
}

// Push pushes one item to stack.
func (s *stackWithMaxValueTrace) Push(x int) {
    s.items = append(s.items, x)
    s.top++

    if len(s.link2NextMaxItem) == 0 {
        s.link2NextMaxItem = append(s.link2NextMaxItem, -1)
        s.maxItemIndex = s.top
    } else {
        m, _ := s.Max()
        if x > m {
            s.link2NextMaxItem = append(s.link2NextMaxItem, s.maxItemIndex)
            s.maxItemIndex = s.top
        }
    }
}

// Pop pops one stack from stack.
func (s *stackWithMaxValueTrace) Pop() (int, error) {
    if s.IsEmpty() {
        return 0, ErrEmptyStack
    }

    if s.maxItemIndex == s.top {
        s.maxItemIndex = s.link2NextMaxItem[s.top]
        s.link2NextMaxItem = s.link2NextMaxItem[:len(s.link2NextMaxItem)-1]
    }

    v := s.items[s.top]
    s.items = s.items[:len(s.items)-1]
    s.top--

    return v, nil
}

// Max tells the max value of current stack.
func (s *stackWithMaxValueTrace) Max() (int, error) {
    if s.IsEmpty() {
        return 0, ErrEmptyStack
    }

    if s.maxItemIndex > -1 {
        return s.items[s.maxItemIndex], nil
    }
    return 0, ErrNoMax
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *stackWithMaxValueTrace) IsEmpty() bool {
    return s.top == -1
}
