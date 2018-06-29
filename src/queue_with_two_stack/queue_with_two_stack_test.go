package queuewithtwostack

import (
	"testing"
)

func TestQueueWithTwoStack(t *testing.T) {
	q := NewQueueWithTwoStack()
	if !q.IsEmpty() {
		t.Error("queue should be empty")
	}

	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	q.PushBack(4)
	q.PushBack(5)

	v, _ := q.PopFront()
	v2, ok := v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 1 {
		t.Errorf("q.PopFront() got %d, want %d", v2, 1)
	}

	v, _ = q.PopFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 2 {
		t.Errorf("q.PopFront() got %d, want %d", v2, 2)
	}

	v, _ = q.PopFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 3 {
		t.Errorf("q.PopFront() got %d, want %d", v2, 3)
	}

	v, _ = q.PopFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 4 {
		t.Errorf("q.PopFront() got %d, want %d", v2, 4)
	}

	v, _ = q.PopFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 5 {
		t.Errorf("q.PopFront() got %d, want %d", v2, 5)
	}
}

func BenchmarkQueuePushBack(b *testing.B) {
	q := NewQueueWithTwoStack()

	for i := 0; i < b.N; i++ {
		q.PushBack(1)
	}
}

func BenchmarkQueuePushBackAndPopFront(b *testing.B) {
	q := NewQueueWithTwoStack()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			q.PushBack(j)
		}
		for j := 0; j < 100; j++ {
			q.PushBack(j)
			q.PopFront()
		}
		for j := 0; j < 100; j++ {
			q.PopFront()
		}
	}
}

func BenchmarkQueuePushBackAndPopFront2(b *testing.B) {
	q := NewQueueWithTwoStack()
	for j := 0; j < 100; j++ {
		q.PushBack(j)
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			q.PushBack(j)
			q.PopFront()
		}
	}
}
