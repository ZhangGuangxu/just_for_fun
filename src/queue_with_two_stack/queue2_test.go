package queuewithtwostack

import (
	"testing"
)

func TestQueue2(t *testing.T) {
	q := newQueue2()
	q.pushBack(1)
	q.pushBack(2)
	q.pushBack(3)
	q.pushBack(4)
	q.pushBack(5)

	v, _ := q.popFront()
	v2, ok := v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 1 {
		t.Errorf("q.popFront() got %d, want %d", v2, 1)
	}

	v, _ = q.popFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 2 {
		t.Errorf("q.popFront() got %d, want %d", v2, 2)
	}

	v, _ = q.popFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 3 {
		t.Errorf("q.popFront() got %d, want %d", v2, 3)
	}

	v, _ = q.popFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 4 {
		t.Errorf("q.popFront() got %d, want %d", v2, 4)
	}

	q.pushBack(101)
	v, _ = q.popFront()
	v2, ok = v.(int)
	if v2 != 5 {
		t.Errorf("q.popFront() got %d, want %d", v2, 5)
	}

	v, _ = q.popFront()
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 101 {
		t.Errorf("q.popFront() got %d, want %d", v2, 101)
	}

	_, err := q.popFront()
	if err == nil {
		t.Error("q.popFront() from empty queue should return error")
	}
}

func BenchmarkQueue2PushBack(b *testing.B) {
	q := newQueue2()
	for i := 0; i < b.N; i++ {
		q.pushBack(1)
	}
}

func BenchmarkQueue2PushBackAndPopFront(b *testing.B) {
	q := newQueue2()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			q.pushBack(j)
		}
		for j := 0; j < 100; j++ {
			q.pushBack(j)
			q.popFront()
		}
		for j := 0; j < 100; j++ {
			q.popFront()
		}
	}
}

func BenchmarkQueue2PushBackAndPopFront2(b *testing.B) {
	q := newQueue2()
	for j := 0; j < 100; j++ {
		q.pushBack(j)
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			q.pushBack(j)
			q.popFront()
		}
	}
}
